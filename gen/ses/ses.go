// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ses provides a client for Amazon Simple Email Service.
package ses

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// SES is a client for Amazon Simple Email Service.
type SES struct {
	client *aws.QueryClient
}

// New returns a new SES client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *SES {
	if client == nil {
		client = http.DefaultClient
	}

	service := "email"
	endpoint, service, region := endpoints.Lookup("email", region)

	return &SES{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-12-01",
		},
	}
}

// DeleteIdentity deletes the specified identity (email address or domain)
// from the list of verified identities. This action is throttled at one
// request per second.
func (c *SES) DeleteIdentity(req *DeleteIdentityRequest) (resp *DeleteIdentityResult, err error) {
	resp = &DeleteIdentityResult{}
	err = c.client.Do("DeleteIdentity", "POST", "/", req, resp)
	return
}

// DeleteVerifiedEmailAddress deletes the specified email address from the
// list of verified addresses. The DeleteVerifiedEmailAddress action is
// deprecated as of the May 15, 2012 release of Domain Verification. The
// DeleteIdentity action is now preferred. This action is throttled at one
// request per second.
func (c *SES) DeleteVerifiedEmailAddress(req *DeleteVerifiedEmailAddressRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVerifiedEmailAddress", "POST", "/", req, nil)
	return
}

// GetIdentityDkimAttributes returns the current status of Easy signing for
// an entity. For domain name identities, this action also returns the
// tokens that are required for Easy signing, and whether Amazon SES has
// successfully verified that these tokens have been published. This action
// takes a list of identities as input and returns the following
// information for each: Whether Easy signing is enabled or disabled. A set
// of tokens that represent the identity. If the identity is an email
// address, the tokens represent the domain of that address. Whether Amazon
// SES has successfully verified the tokens published in the domain's This
// information is only returned for domain name identities, not for email
// addresses. This action is throttled at one request per second. For more
// information about creating DNS records using tokens, go to the Amazon
// SES Developer Guide
func (c *SES) GetIdentityDkimAttributes(req *GetIdentityDkimAttributesRequest) (resp *GetIdentityDkimAttributesResult, err error) {
	resp = &GetIdentityDkimAttributesResult{}
	err = c.client.Do("GetIdentityDkimAttributes", "POST", "/", req, resp)
	return
}

// GetIdentityNotificationAttributes given a list of verified identities
// (email addresses and/or domains), returns a structure describing
// identity notification attributes. This action is throttled at one
// request per second. For more information about using notifications with
// Amazon see the Amazon SES Developer Guide
func (c *SES) GetIdentityNotificationAttributes(req *GetIdentityNotificationAttributesRequest) (resp *GetIdentityNotificationAttributesResult, err error) {
	resp = &GetIdentityNotificationAttributesResult{}
	err = c.client.Do("GetIdentityNotificationAttributes", "POST", "/", req, resp)
	return
}

// GetIdentityVerificationAttributes given a list of identities (email
// addresses and/or domains), returns the verification status and (for
// domain identities) the verification token for each identity. This action
// is throttled at one request per second.
func (c *SES) GetIdentityVerificationAttributes(req *GetIdentityVerificationAttributesRequest) (resp *GetIdentityVerificationAttributesResult, err error) {
	resp = &GetIdentityVerificationAttributesResult{}
	err = c.client.Do("GetIdentityVerificationAttributes", "POST", "/", req, resp)
	return
}

// GetSendQuota returns the user's current sending limits. This action is
// throttled at one request per second.
func (c *SES) GetSendQuota() (resp *GetSendQuotaResult, err error) {
	resp = &GetSendQuotaResult{}
	err = c.client.Do("GetSendQuota", "POST", "/", nil, resp)
	return
}

// GetSendStatistics returns the user's sending statistics. The result is a
// list of data points, representing the last two weeks of sending
// activity. Each data point in the list contains statistics for a
// 15-minute interval. This action is throttled at one request per second.
func (c *SES) GetSendStatistics() (resp *GetSendStatisticsResult, err error) {
	resp = &GetSendStatisticsResult{}
	err = c.client.Do("GetSendStatistics", "POST", "/", nil, resp)
	return
}

// ListIdentities returns a list containing all of the identities (email
// addresses and domains) for a specific AWS Account, regardless of
// verification status. This action is throttled at one request per second.
func (c *SES) ListIdentities(req *ListIdentitiesRequest) (resp *ListIdentitiesResult, err error) {
	resp = &ListIdentitiesResult{}
	err = c.client.Do("ListIdentities", "POST", "/", req, resp)
	return
}

// ListVerifiedEmailAddresses returns a list containing all of the email
// addresses that have been verified. The ListVerifiedEmailAddresses action
// is deprecated as of the May 15, 2012 release of Domain Verification. The
// ListIdentities action is now preferred. This action is throttled at one
// request per second.
func (c *SES) ListVerifiedEmailAddresses() (resp *ListVerifiedEmailAddressesResult, err error) {
	resp = &ListVerifiedEmailAddressesResult{}
	err = c.client.Do("ListVerifiedEmailAddresses", "POST", "/", nil, resp)
	return
}

// SendEmail composes an email message based on input data, and then
// immediately queues the message for sending. You can only send email from
// verified email addresses and domains. If you have not requested
// production access to Amazon you must also verify every recipient email
// address except for the recipients provided by the Amazon SES mailbox
// simulator. For more information, go to the Amazon SES Developer Guide .
// The total size of the message cannot exceed 10 Amazon SES has a limit on
// the total number of recipients per message: The combined number of To:,
// CC: and email addresses cannot exceed 50. If you need to send an email
// message to a larger audience, you can divide your recipient list into
// groups of 50 or fewer, and then call Amazon SES repeatedly to send the
// message to each group. For every message that you send, the total number
// of recipients (To:, CC: and is counted against your sending quota - the
// maximum number of emails you can send in a 24-hour period. For
// information about your sending quota, go to the Amazon SES Developer
// Guide .
func (c *SES) SendEmail(req *SendEmailRequest) (resp *SendEmailResult, err error) {
	resp = &SendEmailResult{}
	err = c.client.Do("SendEmail", "POST", "/", req, resp)
	return
}

// SendRawEmail sends an email message, with header and content specified
// by the client. The SendRawEmail action is useful for sending multipart
// emails. The raw text of the message must comply with Internet email
// standards; otherwise, the message cannot be sent. You can only send
// email from verified email addresses and domains. If you have not
// requested production access to Amazon you must also verify every
// recipient email address except for the recipients provided by the Amazon
// SES mailbox simulator. For more information, go to the Amazon SES
// Developer Guide . The total size of the message cannot exceed 10 MB.
// This includes any attachments that are part of the message. Amazon SES
// has a limit on the total number of recipients per message: The combined
// number of To:, CC: and email addresses cannot exceed 50. If you need to
// send an email message to a larger audience, you can divide your
// recipient list into groups of 50 or fewer, and then call Amazon SES
// repeatedly to send the message to each group. The To:, and headers in
// the raw message can contain a group list. Note that each recipient in a
// group list counts towards the 50-recipient limit. For every message that
// you send, the total number of recipients (To:, CC: and is counted
// against your sending quota - the maximum number of emails you can send
// in a 24-hour period. For information about your sending quota, go to the
// Amazon SES Developer Guide .
func (c *SES) SendRawEmail(req *SendRawEmailRequest) (resp *SendRawEmailResult, err error) {
	resp = &SendRawEmailResult{}
	err = c.client.Do("SendRawEmail", "POST", "/", req, resp)
	return
}

// SetIdentityDkimEnabled enables or disables Easy signing of email sent
// from an identity: If Easy signing is enabled for a domain name identity
// (e.g., example.com ), then Amazon SES will DKIM-sign all email sent by
// addresses under that domain name (e.g., user@example.com If Easy signing
// is enabled for an email address, then Amazon SES will DKIM-sign all
// email sent by that email address. For email addresses (e.g.,
// user@example.com ), you can only enable Easy signing if the
// corresponding domain (e.g., example.com ) has been set up for Easy using
// the AWS Console or the VerifyDomainDkim action. This action is throttled
// at one request per second. For more information about Easy signing, go
// to the Amazon SES Developer Guide
func (c *SES) SetIdentityDkimEnabled(req *SetIdentityDkimEnabledRequest) (resp *SetIdentityDkimEnabledResult, err error) {
	resp = &SetIdentityDkimEnabledResult{}
	err = c.client.Do("SetIdentityDkimEnabled", "POST", "/", req, resp)
	return
}

// SetIdentityFeedbackForwardingEnabled given an identity (email address or
// domain), enables or disables whether Amazon SES forwards bounce and
// complaint notifications as email. Feedback forwarding can only be
// disabled when Amazon Simple Notification Service (Amazon topics are
// specified for both bounces and complaints. This action is throttled at
// one request per second. For more information about using notifications
// with Amazon see the Amazon SES Developer Guide
func (c *SES) SetIdentityFeedbackForwardingEnabled(req *SetIdentityFeedbackForwardingEnabledRequest) (resp *SetIdentityFeedbackForwardingEnabledResult, err error) {
	resp = &SetIdentityFeedbackForwardingEnabledResult{}
	err = c.client.Do("SetIdentityFeedbackForwardingEnabled", "POST", "/", req, resp)
	return
}

// SetIdentityNotificationTopic given an identity (email address or
// domain), sets the Amazon Simple Notification Service (Amazon topic to
// which Amazon SES will publish bounce, complaint, and/or delivery
// notifications for emails sent with that identity as the Source This
// action is throttled at one request per second. For more information
// about feedback notification, see the Amazon SES Developer Guide
func (c *SES) SetIdentityNotificationTopic(req *SetIdentityNotificationTopicRequest) (resp *SetIdentityNotificationTopicResult, err error) {
	resp = &SetIdentityNotificationTopicResult{}
	err = c.client.Do("SetIdentityNotificationTopic", "POST", "/", req, resp)
	return
}

// VerifyDomainDkim returns a set of tokens for a domain. tokens are
// character strings that represent your domain's identity. Using these
// tokens, you will need to create DNS records that point to public keys
// hosted by Amazon Amazon Web Services will eventually detect that you
// have updated your DNS records; this detection process may take up to 72
// hours. Upon successful detection, Amazon SES will be able to DKIM-sign
// email originating from that domain. This action is throttled at one
// request per second. To enable or disable Easy signing for a domain, use
// the SetIdentityDkimEnabled action. For more information about creating
// DNS records using tokens, go to the Amazon SES Developer Guide
func (c *SES) VerifyDomainDkim(req *VerifyDomainDkimRequest) (resp *VerifyDomainDkimResult, err error) {
	resp = &VerifyDomainDkimResult{}
	err = c.client.Do("VerifyDomainDkim", "POST", "/", req, resp)
	return
}

// VerifyDomainIdentity verifies a domain. This action is throttled at one
// request per second.
func (c *SES) VerifyDomainIdentity(req *VerifyDomainIdentityRequest) (resp *VerifyDomainIdentityResult, err error) {
	resp = &VerifyDomainIdentityResult{}
	err = c.client.Do("VerifyDomainIdentity", "POST", "/", req, resp)
	return
}

// VerifyEmailAddress verifies an email address. This action causes a
// confirmation email message to be sent to the specified address. The
// VerifyEmailAddress action is deprecated as of the May 15, 2012 release
// of Domain Verification. The VerifyEmailIdentity action is now preferred.
// This action is throttled at one request per second.
func (c *SES) VerifyEmailAddress(req *VerifyEmailAddressRequest) (err error) {
	// NRE
	err = c.client.Do("VerifyEmailAddress", "POST", "/", req, nil)
	return
}

// VerifyEmailIdentity verifies an email address. This action causes a
// confirmation email message to be sent to the specified address. This
// action is throttled at one request per second.
func (c *SES) VerifyEmailIdentity(req *VerifyEmailIdentityRequest) (resp *VerifyEmailIdentityResult, err error) {
	resp = &VerifyEmailIdentityResult{}
	err = c.client.Do("VerifyEmailIdentity", "POST", "/", req, resp)
	return
}

// Body is undocumented.
type Body struct {
	HTML *Content `query:"Html" xml:"Html"`
	Text *Content `query:"Text" xml:"Text"`
}

// Content is undocumented.
type Content struct {
	Charset aws.StringValue `query:"Charset" xml:"Charset"`
	Data    aws.StringValue `query:"Data" xml:"Data"`
}

// DeleteIdentityRequest is undocumented.
type DeleteIdentityRequest struct {
	Identity aws.StringValue `query:"Identity" xml:"Identity"`
}

// DeleteIdentityResponse is undocumented.
type DeleteIdentityResponse struct {
}

// DeleteVerifiedEmailAddressRequest is undocumented.
type DeleteVerifiedEmailAddressRequest struct {
	EmailAddress aws.StringValue `query:"EmailAddress" xml:"EmailAddress"`
}

// Destination is undocumented.
type Destination struct {
	BCCAddresses []string `query:"BccAddresses.member" xml:"BccAddresses>member"`
	CCAddresses  []string `query:"CcAddresses.member" xml:"CcAddresses>member"`
	ToAddresses  []string `query:"ToAddresses.member" xml:"ToAddresses>member"`
}

// GetIdentityDkimAttributesRequest is undocumented.
type GetIdentityDkimAttributesRequest struct {
	Identities []string `query:"Identities.member" xml:"Identities>member"`
}

// GetIdentityDkimAttributesResponse is undocumented.
type GetIdentityDkimAttributesResponse struct {
	DkimAttributes map[string]IdentityDkimAttributes `query:"DkimAttributes" xml:"GetIdentityDkimAttributesResult>DkimAttributes"`
}

// GetIdentityNotificationAttributesRequest is undocumented.
type GetIdentityNotificationAttributesRequest struct {
	Identities []string `query:"Identities.member" xml:"Identities>member"`
}

// GetIdentityNotificationAttributesResponse is undocumented.
type GetIdentityNotificationAttributesResponse struct {
	NotificationAttributes map[string]IdentityNotificationAttributes `query:"NotificationAttributes" xml:"GetIdentityNotificationAttributesResult>NotificationAttributes"`
}

// GetIdentityVerificationAttributesRequest is undocumented.
type GetIdentityVerificationAttributesRequest struct {
	Identities []string `query:"Identities.member" xml:"Identities>member"`
}

// GetIdentityVerificationAttributesResponse is undocumented.
type GetIdentityVerificationAttributesResponse struct {
	VerificationAttributes map[string]IdentityVerificationAttributes `query:"VerificationAttributes" xml:"GetIdentityVerificationAttributesResult>VerificationAttributes"`
}

// GetSendQuotaResponse is undocumented.
type GetSendQuotaResponse struct {
	Max24HourSend   aws.DoubleValue `query:"Max24HourSend" xml:"GetSendQuotaResult>Max24HourSend"`
	MaxSendRate     aws.DoubleValue `query:"MaxSendRate" xml:"GetSendQuotaResult>MaxSendRate"`
	SentLast24Hours aws.DoubleValue `query:"SentLast24Hours" xml:"GetSendQuotaResult>SentLast24Hours"`
}

// GetSendStatisticsResponse is undocumented.
type GetSendStatisticsResponse struct {
	SendDataPoints []SendDataPoint `query:"SendDataPoints.member" xml:"GetSendStatisticsResult>SendDataPoints>member"`
}

// IdentityDkimAttributes is undocumented.
type IdentityDkimAttributes struct {
	DkimEnabled            aws.BooleanValue `query:"DkimEnabled" xml:"DkimEnabled"`
	DkimTokens             []string         `query:"DkimTokens.member" xml:"DkimTokens>member"`
	DkimVerificationStatus aws.StringValue  `query:"DkimVerificationStatus" xml:"DkimVerificationStatus"`
}

// IdentityNotificationAttributes is undocumented.
type IdentityNotificationAttributes struct {
	BounceTopic       aws.StringValue  `query:"BounceTopic" xml:"BounceTopic"`
	ComplaintTopic    aws.StringValue  `query:"ComplaintTopic" xml:"ComplaintTopic"`
	DeliveryTopic     aws.StringValue  `query:"DeliveryTopic" xml:"DeliveryTopic"`
	ForwardingEnabled aws.BooleanValue `query:"ForwardingEnabled" xml:"ForwardingEnabled"`
}

// Possible values for SES.
const (
	IdentityTypeDomain       = "Domain"
	IdentityTypeEmailAddress = "EmailAddress"
)

// IdentityVerificationAttributes is undocumented.
type IdentityVerificationAttributes struct {
	VerificationStatus aws.StringValue `query:"VerificationStatus" xml:"VerificationStatus"`
	VerificationToken  aws.StringValue `query:"VerificationToken" xml:"VerificationToken"`
}

// ListIdentitiesRequest is undocumented.
type ListIdentitiesRequest struct {
	IdentityType aws.StringValue  `query:"IdentityType" xml:"IdentityType"`
	MaxItems     aws.IntegerValue `query:"MaxItems" xml:"MaxItems"`
	NextToken    aws.StringValue  `query:"NextToken" xml:"NextToken"`
}

// ListIdentitiesResponse is undocumented.
type ListIdentitiesResponse struct {
	Identities []string        `query:"Identities.member" xml:"ListIdentitiesResult>Identities>member"`
	NextToken  aws.StringValue `query:"NextToken" xml:"ListIdentitiesResult>NextToken"`
}

// ListVerifiedEmailAddressesResponse is undocumented.
type ListVerifiedEmailAddressesResponse struct {
	VerifiedEmailAddresses []string `query:"VerifiedEmailAddresses.member" xml:"ListVerifiedEmailAddressesResult>VerifiedEmailAddresses>member"`
}

// Message is undocumented.
type Message struct {
	Body    *Body    `query:"Body" xml:"Body"`
	Subject *Content `query:"Subject" xml:"Subject"`
}

// Possible values for SES.
const (
	NotificationTypeBounce    = "Bounce"
	NotificationTypeComplaint = "Complaint"
	NotificationTypeDelivery  = "Delivery"
)

// RawMessage is undocumented.
type RawMessage struct {
	Data []byte `query:"Data" xml:"Data"`
}

// SendDataPoint is undocumented.
type SendDataPoint struct {
	Bounces          aws.LongValue `query:"Bounces" xml:"Bounces"`
	Complaints       aws.LongValue `query:"Complaints" xml:"Complaints"`
	DeliveryAttempts aws.LongValue `query:"DeliveryAttempts" xml:"DeliveryAttempts"`
	Rejects          aws.LongValue `query:"Rejects" xml:"Rejects"`
	Timestamp        time.Time     `query:"Timestamp" xml:"Timestamp"`
}

// SendEmailRequest is undocumented.
type SendEmailRequest struct {
	Destination      *Destination    `query:"Destination" xml:"Destination"`
	Message          *Message        `query:"Message" xml:"Message"`
	ReplyToAddresses []string        `query:"ReplyToAddresses.member" xml:"ReplyToAddresses>member"`
	ReturnPath       aws.StringValue `query:"ReturnPath" xml:"ReturnPath"`
	Source           aws.StringValue `query:"Source" xml:"Source"`
}

// SendEmailResponse is undocumented.
type SendEmailResponse struct {
	MessageID aws.StringValue `query:"MessageId" xml:"SendEmailResult>MessageId"`
}

// SendRawEmailRequest is undocumented.
type SendRawEmailRequest struct {
	Destinations []string        `query:"Destinations.member" xml:"Destinations>member"`
	RawMessage   *RawMessage     `query:"RawMessage" xml:"RawMessage"`
	Source       aws.StringValue `query:"Source" xml:"Source"`
}

// SendRawEmailResponse is undocumented.
type SendRawEmailResponse struct {
	MessageID aws.StringValue `query:"MessageId" xml:"SendRawEmailResult>MessageId"`
}

// SetIdentityDkimEnabledRequest is undocumented.
type SetIdentityDkimEnabledRequest struct {
	DkimEnabled aws.BooleanValue `query:"DkimEnabled" xml:"DkimEnabled"`
	Identity    aws.StringValue  `query:"Identity" xml:"Identity"`
}

// SetIdentityDkimEnabledResponse is undocumented.
type SetIdentityDkimEnabledResponse struct {
}

// SetIdentityFeedbackForwardingEnabledRequest is undocumented.
type SetIdentityFeedbackForwardingEnabledRequest struct {
	ForwardingEnabled aws.BooleanValue `query:"ForwardingEnabled" xml:"ForwardingEnabled"`
	Identity          aws.StringValue  `query:"Identity" xml:"Identity"`
}

// SetIdentityFeedbackForwardingEnabledResponse is undocumented.
type SetIdentityFeedbackForwardingEnabledResponse struct {
}

// SetIdentityNotificationTopicRequest is undocumented.
type SetIdentityNotificationTopicRequest struct {
	Identity         aws.StringValue `query:"Identity" xml:"Identity"`
	NotificationType aws.StringValue `query:"NotificationType" xml:"NotificationType"`
	SNSTopic         aws.StringValue `query:"SnsTopic" xml:"SnsTopic"`
}

// SetIdentityNotificationTopicResponse is undocumented.
type SetIdentityNotificationTopicResponse struct {
}

// Possible values for SES.
const (
	VerificationStatusFailed           = "Failed"
	VerificationStatusNotStarted       = "NotStarted"
	VerificationStatusPending          = "Pending"
	VerificationStatusSuccess          = "Success"
	VerificationStatusTemporaryFailure = "TemporaryFailure"
)

// VerifyDomainDkimRequest is undocumented.
type VerifyDomainDkimRequest struct {
	Domain aws.StringValue `query:"Domain" xml:"Domain"`
}

// VerifyDomainDkimResponse is undocumented.
type VerifyDomainDkimResponse struct {
	DkimTokens []string `query:"DkimTokens.member" xml:"VerifyDomainDkimResult>DkimTokens>member"`
}

// VerifyDomainIdentityRequest is undocumented.
type VerifyDomainIdentityRequest struct {
	Domain aws.StringValue `query:"Domain" xml:"Domain"`
}

// VerifyDomainIdentityResponse is undocumented.
type VerifyDomainIdentityResponse struct {
	VerificationToken aws.StringValue `query:"VerificationToken" xml:"VerifyDomainIdentityResult>VerificationToken"`
}

// VerifyEmailAddressRequest is undocumented.
type VerifyEmailAddressRequest struct {
	EmailAddress aws.StringValue `query:"EmailAddress" xml:"EmailAddress"`
}

// VerifyEmailIdentityRequest is undocumented.
type VerifyEmailIdentityRequest struct {
	EmailAddress aws.StringValue `query:"EmailAddress" xml:"EmailAddress"`
}

// VerifyEmailIdentityResponse is undocumented.
type VerifyEmailIdentityResponse struct {
}

// DeleteIdentityResult is a wrapper for DeleteIdentityResponse.
type DeleteIdentityResult struct {
}

// GetIdentityDkimAttributesResult is a wrapper for GetIdentityDkimAttributesResponse.
type GetIdentityDkimAttributesResult struct {
	DkimAttributes map[string]IdentityDkimAttributes `query:"DkimAttributes" xml:"GetIdentityDkimAttributesResult>DkimAttributes"`
}

// GetIdentityNotificationAttributesResult is a wrapper for GetIdentityNotificationAttributesResponse.
type GetIdentityNotificationAttributesResult struct {
	NotificationAttributes map[string]IdentityNotificationAttributes `query:"NotificationAttributes" xml:"GetIdentityNotificationAttributesResult>NotificationAttributes"`
}

// GetIdentityVerificationAttributesResult is a wrapper for GetIdentityVerificationAttributesResponse.
type GetIdentityVerificationAttributesResult struct {
	VerificationAttributes map[string]IdentityVerificationAttributes `query:"VerificationAttributes" xml:"GetIdentityVerificationAttributesResult>VerificationAttributes"`
}

// GetSendQuotaResult is a wrapper for GetSendQuotaResponse.
type GetSendQuotaResult struct {
	Max24HourSend   aws.DoubleValue `query:"Max24HourSend" xml:"GetSendQuotaResult>Max24HourSend"`
	MaxSendRate     aws.DoubleValue `query:"MaxSendRate" xml:"GetSendQuotaResult>MaxSendRate"`
	SentLast24Hours aws.DoubleValue `query:"SentLast24Hours" xml:"GetSendQuotaResult>SentLast24Hours"`
}

// GetSendStatisticsResult is a wrapper for GetSendStatisticsResponse.
type GetSendStatisticsResult struct {
	SendDataPoints []SendDataPoint `query:"SendDataPoints.member" xml:"GetSendStatisticsResult>SendDataPoints>member"`
}

// ListIdentitiesResult is a wrapper for ListIdentitiesResponse.
type ListIdentitiesResult struct {
	Identities []string        `query:"Identities.member" xml:"ListIdentitiesResult>Identities>member"`
	NextToken  aws.StringValue `query:"NextToken" xml:"ListIdentitiesResult>NextToken"`
}

// ListVerifiedEmailAddressesResult is a wrapper for ListVerifiedEmailAddressesResponse.
type ListVerifiedEmailAddressesResult struct {
	VerifiedEmailAddresses []string `query:"VerifiedEmailAddresses.member" xml:"ListVerifiedEmailAddressesResult>VerifiedEmailAddresses>member"`
}

// SendEmailResult is a wrapper for SendEmailResponse.
type SendEmailResult struct {
	MessageID aws.StringValue `query:"MessageId" xml:"SendEmailResult>MessageId"`
}

// SendRawEmailResult is a wrapper for SendRawEmailResponse.
type SendRawEmailResult struct {
	MessageID aws.StringValue `query:"MessageId" xml:"SendRawEmailResult>MessageId"`
}

// SetIdentityDkimEnabledResult is a wrapper for SetIdentityDkimEnabledResponse.
type SetIdentityDkimEnabledResult struct {
}

// SetIdentityFeedbackForwardingEnabledResult is a wrapper for SetIdentityFeedbackForwardingEnabledResponse.
type SetIdentityFeedbackForwardingEnabledResult struct {
}

// SetIdentityNotificationTopicResult is a wrapper for SetIdentityNotificationTopicResponse.
type SetIdentityNotificationTopicResult struct {
}

// VerifyDomainDkimResult is a wrapper for VerifyDomainDkimResponse.
type VerifyDomainDkimResult struct {
	DkimTokens []string `query:"DkimTokens.member" xml:"VerifyDomainDkimResult>DkimTokens>member"`
}

// VerifyDomainIdentityResult is a wrapper for VerifyDomainIdentityResponse.
type VerifyDomainIdentityResult struct {
	VerificationToken aws.StringValue `query:"VerificationToken" xml:"VerifyDomainIdentityResult>VerificationToken"`
}

// VerifyEmailIdentityResult is a wrapper for VerifyEmailIdentityResponse.
type VerifyEmailIdentityResult struct {
}

// avoid errors if the packages aren't referenced
var _ time.Time
