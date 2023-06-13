package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/rumodel"
	"github.com/NubeIO/rubix-ui/backend/utils/urls"
)

func (inst *Client) AddProducer(hostIDName string, body *model.Producer) (*model.Producer, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Producer{}).
		SetBody(body).
		Post("/proxy/ros/api/producers"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *Client) GetProducers(hostIDName string) ([]model.Producer, error) {
	url := fmt.Sprintf("/proxy/ros/api/producers")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Producer
	out = *resp.Result().(*[]model.Producer)
	return out, nil
}

func (inst *Client) GetProducer(hostIDName, uuid string, withWriterClones bool) (*model.Producer, error) {
	url := fmt.Sprintf("/proxy/ros/api/producers/%s", uuid)
	if withWriterClones {
		url = urls.AttachQueryParams(url, "with_writer_clones=true")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *Client) GetProducerByThingUUID(hostIDName, thingUUID string) (*model.Producer, error) {
	url := fmt.Sprintf("/proxy/ros/api/producers/one/args?producer_thing_uuid=%s", thingUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *Client) GetProducerOneArg(hostIDName, arg string) (*model.Producer, error) {
	url := fmt.Sprintf("/proxy/ros/api/producers/one/args?%s", arg)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *Client) EditProducer(hostIDName, uuid string, body *model.Producer) (*model.Producer, error) {
	url := fmt.Sprintf("/proxy/ros/api/producers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Producer{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *Client) DeleteProducer(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/proxy/ros/api/producers/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) SyncProducers(hostIDName, streamUUID string) (*[]rumodel.SyncModel, error) {
	url := fmt.Sprintf(
		"/proxy/ros/api/streams/%s/sync/producers?with_writer_clones=true", streamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]rumodel.SyncModel{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]rumodel.SyncModel), nil
}
