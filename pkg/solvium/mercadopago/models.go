package mercadopago

type ErrorResponse struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	ErrorCode string `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

type PaymentIdentification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Payer struct {
	Name           string                `json:"name"`
	Surname        string                `json:"surname"`
	Email          string                `json:"email"`
	Identification PaymentIdentification `json:"identification"`
}

type Item struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type BackURLs struct {
	Success string `json:"success,omitempty"`
	Pending string `json:"pending,omitempty"`
	Failure string `json:"failure,omitempty"`
}

type PaymentRequest struct {
	ExternalReference string    `json:"external_reference,omitempty"`
	Items             []Item    `json:"items"`
	Payer             Payer     `json:"payer"`
	BackURLs          *BackURLs `json:"back_urls,omitempty"`
	NotificationURL   string    `json:"notification_url,omitempty"`
	AutoReturn        string    `json:"auto_return,omitempty"`
	ExpirationDateTo  string    `json:"expiration_date_to,omitempty"`
}

type PaymentResponse struct {
	ID                string    `json:"id"`
	InitPoint         string    `json:"init_point"`
	SandboxInitPoint  string    `json:"sandbox_init_point"`
	ExternalReference string    `json:"external_reference"`
	Items             []Item    `json:"items"`
	Payer             Payer     `json:"payer"`
	BackURLs          *BackURLs `json:"back_urls"`
	NotificationURL   string    `json:"notification_url"`
	CreationDate      string    `json:"date_created"`
}

type PaymentConsultResponse struct {
	ID                string  `json:"id"`
	DateCreated       string  `json:"date_created"`
	DateApproved      string  `json:"date_approved"`
	DateLastUpdated   string  `json:"date_last_updated"`
	DateOfExpiration  string  `json:"date_of_expiration"`
	MoneyReleaseDate  string  `json:"money_release_date"`
	OperationType     string  `json:"operation_type"`
	IssuerId          string  `json:"issuer_id"`
	PaymentMethodId   string  `json:"payment_method_id"`
	PaymentTypeId     string  `json:"payment_type_id"`
	Status            string  `json:"status"`
	StatusDetail      string  `json:"status_detail"`
	CurrencyId        string  `json:"currency_id"`
	Description       string  `json:"description"`
	LiveMode          bool    `json:"live_mode"`
	ExternalReference string  `json:"external_reference"`
	TransactionAmount float64 `json:"transaction_amount"`
}

type PaymentSearchParams map[string]string

type PaymentSearchResult struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type PaymentSearchResponse struct {
	Paging struct {
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"paging"`
	Results []PaymentSearchResult `json:"results"`
}

type IdentificationType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}

type PaymentMethod struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	PaymentTypeID    string   `json:"payment_type_id"`
	Status           string   `json:"status"`
	SecureThumbnail  string   `json:"secure_thumbnail"`
	Thumbnail        string   `json:"thumbnail"`
	DeferredCapture  string   `json:"deferred_capture"`
	Description      string   `json:"description"`
	MinAllowedAmount float64  `json:"min_allowed_amount"`
	MaxAllowedAmount float64  `json:"max_allowed_amount"`
	ProcessingModes  []string `json:"processing_modes"`
}
