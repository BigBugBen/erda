package bundle

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle/apierrors"
	"github.com/erda-project/erda/pkg/httputil"
)

func (b *Bundle) CreateNoticeRequest(req *apistructs.NoticeCreateRequest,
	orgID uint64) (*apistructs.NoticeCreateResponse, error) {
	cmdbURL, err := b.urls.CMDB()
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}

	var buf bytes.Buffer
	httpClient := b.hc
	resp, err := httpClient.Post(cmdbURL).Path("/api/notices").
		Header(httputil.OrgHeader, fmt.Sprintf("%d", orgID)).
		Header(httputil.UserHeader, "1").
		JSONBody(&req).Do().Body(&buf)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !resp.IsOK() {
		return nil, apierrors.ErrInvoke.InternalError(
			fmt.Errorf("failed to create notice, status code: %d, body: %v",
				resp.StatusCode(),
				buf.String(),
			))
	}
	// TODO: delete the Notice when unmarshal error
	var ncresp apistructs.NoticeCreateResponse
	err = json.Unmarshal(buf.Bytes(), &ncresp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	return &ncresp, nil
}

func (b *Bundle) PublishORUnPublishNotice(orgID uint64, noticeID uint64, publishType string) error {
	cmdbURL, err := b.urls.CMDB()
	if err != nil {
		return apierrors.ErrInvoke.InternalError(err)
	}

	var buf bytes.Buffer
	resp, err := b.hc.Put(cmdbURL).Path(fmt.Sprintf("/api/notices/%d/actions/%s", noticeID, publishType)).
		Header(httputil.OrgHeader, fmt.Sprintf("%d", orgID)).
		Header(httputil.UserHeader, "1").
		Do().
		Body(&buf)
	if err != nil {
		return apierrors.ErrInvoke.InternalError(err)
	}

	if !resp.IsOK() {
		return apierrors.ErrInvoke.InternalError(
			fmt.Errorf("failed to %s notice, status code: %d, body: %v",
				publishType,
				resp.StatusCode(),
				buf.String(),
			))
	}
	return nil
}

func (b *Bundle) ListNoticeByOrgID(orgID uint64) (*apistructs.NoticeListResponse, error) {
	cmdbURL, err := b.urls.CMDB()
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}

	var buf bytes.Buffer
	resp, err := b.hc.Get(cmdbURL).Path("/api/notices").
		Header(httputil.OrgHeader, fmt.Sprintf("%d", orgID)).
		Header(httputil.UserHeader, "1").
		Do().
		Body(&buf)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}

	if !resp.IsOK() {
		return nil, apierrors.ErrInvoke.InternalError(
			fmt.Errorf("failed to list notice, status code: %d, body: %v",
				resp.StatusCode(),
				buf.String(),
			))
	}

	var notelist apistructs.NoticeListResponse
	err = json.Unmarshal(buf.Bytes(), &notelist)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	return &notelist, nil
}
