package servicedesk

import (
	"io"

	"github.com/andygrunwald/go-jira/v2/cloud/model"
)

type TemporaryFile struct {
	Name string
	File io.Reader
}

type AttachedTemporaryFile struct {
	TemporaryAttachments []TemporaryAttachment `json:"temporaryAttachments,omitempty" structs:"temporaryAttachments,omitempty"`
}

type TemporaryAttachment struct {
	TemporaryAttachmentID string `json:"temporaryAttachmentId,omitempty" structs:"temporaryAttachmentId,omitempty"`
	FileName              string `json:"fileName,omitempty" structs:"fileName,omitempty"`
}

type AttachmentDTO struct {
	Filename string            `json:"filename,omitempty" structs:"filename,omitempty"`
	Author   UserDTO           `json:"author,omitempty" structs:"author,omitempty"`
	Created  model.DateDTO     `json:"created,omitempty" structs:"created,omitempty"`
	Size     int64             `json:"size,omitempty" structs:"size,omitempty"`
	MIMEType string            `json:"mimeType,omitempty" structs:"mimeType,omitempty"`
	Links    AttachmentLinkDTO `json:"_links,omitempty" structs:"_links,omitempty"`
}

type AttachmentLinkDTO struct {
	Self      string `json:"self,omitempty" structs:"self,omitempty"`
	JIRARest  string `json:"jiraRest,omitempty" structs:"jiraRest,omitempty"`
	Content   string `json:"content,omitempty" structs:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty" structs:"thumbnail,omitempty"`
}

type AttachmentCreateResultDTO struct {
	Comment     CommentDTO                     `json:"comment,omitempty" structs:"comment,omitempty"`
	Attachments model.PagedDTOT[AttachmentDTO] `json:"attachments,omitempty" structs:"attachments,omitempty"`
}

type CreateRequestAttachment struct {
	TemporaryAttachmentIDs []string              `json:"temporaryAttachmentIds,omitempty" structs:"temporaryAttachmentIds,omitempty"`
	AdditionalComment      *AdditionalCommentDTO `json:"additionalComment,omitempty" structs:"additionalComment,omitempty"`
	Public                 bool                  `json:"public" structs:"public"`
}

type AdditionalCommentDTO struct {
	Body string `json:"body,omitempty" structs:"body,omitempty"`
}

type CommentDTO struct {
	ID           string
	Body         string                         `json:"body,omitempty" structs:"body,omitempty"`
	RenderedBody *RenderedValueDTO              `json:"renderedBody,omitempty" structs:"renderedBody,omitempty"`
	Author       UserDTO                        `json:"author,omitempty" structs:"author,omitempty"`
	Created      model.DateDTO                  `json:"created,omitempty" structs:"created,omitempty"`
	Attachments  model.PagedDTOT[AttachmentDTO] `json:"attachments,omitempty" structs:"attachments,omitempty"`
	Public       bool                           `json:"public,omitempty" structs:"public,omitempty"`
	Expands      []string                       `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links        *model.SelfLink                `json:"_links,omitempty" structs:"_links,omitempty"`
}
