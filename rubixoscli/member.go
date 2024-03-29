package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

type UpdateMemberPasswordBody struct {
	NewPassword string `json:"new_password"`
}

func (inst *Client) GetMembers() ([]model.Member, error) {
	url := "/api/members"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Member{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Member
	out = *resp.Result().(*[]model.Member)
	return out, nil
}

func (inst *Client) AddMember(body *model.Member) (*model.Member, error) {
	url := "/api/apps/members"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Member{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Member)
	return out, nil
}

func (inst *Client) UpdateMember(memberUUID string, body *model.Member) (*model.Member, error) {
	url := fmt.Sprintf("/api/members/%s", memberUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Member{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Member)
	return out, nil
}

func (inst *Client) DeleteMember(memberUUID string) (bool, error) {
	url := fmt.Sprintf("/api/members/%s", memberUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) BulkDeleteMember(memberUUID []string) (bool, error) {
	for _, aUUID := range memberUUID {
		url := fmt.Sprintf("/api/members/%s", aUUID)
		resp, err := nresty.FormatRestyResponse(inst.Rest.R().
			Delete(url))
		if err != nil || resp.String() != "true" {
			return false, err
		}
	}
	return true, nil
}

func (inst *Client) VerifyMember(memberName string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/members/username/%s/verify", memberName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) UpdateMemberPassword(memberUUID string, password string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/members/%s/change-password", memberUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(UpdateMemberPasswordBody{
			NewPassword: password,
		}).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}
