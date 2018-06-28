package template

import "time"

const (
	TemplateTypeAirlineBoardingpass TemplateType = "airline_boardingpass"
	TemplateTypeAirlineCheckin      TemplateType = "airline_checkin"
	TemplateTypeAirlineItinerary    TemplateType = "airline_itinerary"
	TemplateTypeAirlineUpdate       TemplateType = "airline_update"
)

type AirlineBaseTemplate struct {
	TemplateBase
	IntroMessage string `json:"intro_message"`
	Locale       string `json:"locale"`
}

type AirlineBoardingpassTemplate struct {
	AirlineBaseTemplate
	BoardingPass []BoardingPass `json:"boarding_pass"`
}

func (AirlineBoardingpassTemplate) Type() TemplateType {
	return TemplateTypeAirlineBoardingpass
}

type BoardingPass struct {
	PassengerName        string     `json:"passenger_name"`
	PnrNumber            string     `json:"pnr_number"`
	Seat                 string     `json:"seat"`
	LogoImageURL         string     `json:"logo_image_url"`
	HeaderImageURL       string     `json:"header_image_url"`
	QrCode               string     `json:"qr_code"`
	AboveBarCodeImageURL string     `json:"above_bar_image_url"`
	AuxiliaryFields      []Field    `json:"auxiliary_fields"`
	SecondaryFields      []Field    `json:"secondary_fields"`
	FlightInfo           FlightInfo `json:"flight_info"`
}

type Field struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type FlightInfo struct {
	ConnectionID     string         `json:"connection_id,omitempty"`
	SegmentID        string         `json:"segment_id,omitempty"`
	FlightNumber     string         `json:"flight_number"`
	AircraftType     string         `json:"aircraft_type,omitempty"`
	TravelClass      string         `json:"travel_class,omitempty"`
	DepartureAirport Airport        `json:"departure_airport"`
	ArrivalAirport   Airport        `json:"arrival_airport"`
	FlightSchedule   FlightSchedule `json:"flight_schedule"`
}

type Airport struct {
	AirportCode string `json:"airport_code"`
	City        string `json:"city"`
	Terminal    string `json:"terminal,omitempty"`
	Gate        string `json:"gate,omitempty"`
}

type FlightSchedule struct {
	BoardingTime  time.Time `json:"boarding_time,omitempty"`
	DepartureTime time.Time `json:"departure_time"`
	Arrivaltime   time.Time `json:"arrival_time"`
}

type AirlineCheckinTempate struct {
	AirlineBaseTemplate
	PnrNumber  string       `json:"pnr_number"`
	CheckinURL string       `json:"checkin_url"`
	FlightInfo []FlightInfo `json:"flight_info"`
}

func (AirlineCheckinTempate) Type() TemplateType {
	return TemplateTypeAirlineCheckin
}

type InteraryTemplate struct {
	AirlineBaseTemplate
	PnrNumber            string                 `json:"pnr_number"`
	PassengerInfo        []PassengerInfo        `json:"passenger_info"`
	FlightInfo           []FlightInfo           `json:"flight_info"`
	PassengerSegmentInfo []PassengerSegmentInfo `json:"passenger_segment_info"`
	PriceInfo            []PriceInfo            `json:"price_info"`
	BasePrice            string                 `json:"base_price"`
	Tax                  string                 `json:"tax"`
	TotalPrice           string                 `json:"total_price"`
	Currency             string                 `json:"currency"`
}

func (InteraryTemplate) Type() TemplateType {
	return TemplateTypeAirlineItinerary
}

type PassengerInfo struct {
	Name         string `json:"name"`
	TicketNumber string `json:"ticket_number"`
	PassengerID  string `json:"passenger_id"`
}

type PassengerSegmentInfo struct {
	SegmentID   string        `json:"segment_id"`
	PassengerID string        `json:"passenger_id"`
	Seat        string        `json:"seat"`
	SeatType    string        `json:"seat_type"`
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
}

type ProductInfo struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type PriceInfo struct {
	Title    string `json:"title"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
