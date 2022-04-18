package adapters

import (
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
)

func ToContact(contact *domain.UserContact) *protobuf.UserContact {
	if contact == nil {
		return nil
	}

	return &protobuf.UserContact{
		ContactId:   contact.ContactID,
		UserId:      contact.UserID,
		ContactName: contact.ContactName,
		Contact:     contact.Contact,
	}
}

func ToGetUserContactRequest(req *protobuf.GetUserContactRequest) *domain.GetUserContactRequest {
	if req == nil {
		return nil
	}

	return &domain.GetUserContactRequest{ContactID: req.ContactId}
}

func ToCreateUserContactRequest(req *protobuf.CreateUserContactRequest) *domain.CreateUserContactRequest {
	if req == nil {
		return nil
	}

	return &domain.CreateUserContactRequest{
		UserID:      req.UserId,
		ContactName: req.ContactName,
		Contact:     req.Contact,
	}
}

func ToGetUserContactResponse(contact *domain.UserContact) *protobuf.GetUserContactResponse {
	return &protobuf.GetUserContactResponse{Contact: ToContact(contact)}
}

func ToUpdateUserContactRequest(req *protobuf.UpdateUserContactRequest) *domain.UpdateUserContactRequest {
	if req == nil {
		return nil
	}

	return &domain.UpdateUserContactRequest{
		ContactID:   req.ContactId,
		ContactName: req.ContactName,
		Contact:     req.Contact,
	}
}

func ToDeleteUserContactRequest(req *protobuf.DeleteUserContactRequest) *domain.DeleteUserContactRequest {
	if req == nil {
		return nil
	}

	return &domain.DeleteUserContactRequest{
		ContactID: req.ContactId,
	}
}
