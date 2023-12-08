package plient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/guisecreator/plient/models"
)

const baseURL = "https://api.pinterest.com/v5"

type Client struct {
	httpClient         *http.Client
	accessToken        string
	clientID           string
	clientSecret       string
	baseURL            string
	DefaultContentType string
}

func NewClient(accessToken, clientSecret, clientID string) (*Client, error) {
	if accessToken == "" {
		return nil, errors.New("accessToken is empty")
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		accessToken:        accessToken,
		clientID:           clientID,
		clientSecret:       clientSecret,
		baseURL:            baseURL,
		DefaultContentType: "application/json",
	}, nil
}

func (c *Client) NewRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Content-Type", c.DefaultContentType)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 200 {
		defer response.Body.Close()

		errorRes, errStatus := handleWrongStatusCode(response)
		if errStatus != nil {
			return nil, errStatus
		}

		return nil, errors.New(fmt.Sprintf(
			"Error response: ErrorCode: %d ErrorMessage: %s",
			errorRes.Code,
			errorRes.Message,
		))
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func handleWrongStatusCode(res *http.Response) (models.ErrorResponse, error) {
	errorRes := models.ErrorResponse{}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return errorRes, errors.New("unable to read response body while handleWrongStatus code")
	}

	err = json.Unmarshal(bytes, &errorRes)
	if err != nil {
		return errorRes, errors.New("unable to unmarshal response body while handleWrongStatus code")
	}

	return errorRes, nil
}

func (c *Client) GetPinById(pinId string) (*models.PinsData, error) {
	url := "/pins/"

	if len(pinId) > 0 {
		url += "&image=" + pinId
	}

	responseBytes, err := c.NewRequest(url + pinId + "/")
	if err != nil {
		return nil, err
	}

	var pin models.PinsData
	err = json.Unmarshal(responseBytes, &pin)
	if err != nil {
		return nil, err
	}

	return &pin, nil
}

func (c *Client) GetPinsById(pinId string) ([]models.PinData, error) {
	pins, err := c.GetPinById(pinId)
	if err != nil {
		return nil, err
	}

	var responseBytes []models.PinData

	responseBytes = append(responseBytes, pins.Items...)

	for len(pins.Items) > 0 {
		pins, err = c.GetPinById(pins.Items[len(pins.Items)-1].Id)
		if err != nil {
			return nil, err
		}

		responseBytes = append(responseBytes, pins.Items...)
	}

	return responseBytes, nil
}

func (c *Client) GetBoard(bookmark string) (*models.BoardsData, error) {
	url := "/boards/?page_size=100"
	if len(bookmark) > 0 {
		url += "&bookmark=" + bookmark
	}

	bytes, err := c.NewRequest(url)
	if err != nil {
		return nil, err
	}

	var board = new(models.BoardsData)
	unmarshalErr := json.Unmarshal(bytes, &board)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return board, nil
}

func (c *Client) GetBoards() ([]models.Board, error) {
	var resultBoards []models.Board

	boards, err := c.GetBoard("")
	if err != nil {
		return nil, err
	}

	resultBoards = append(resultBoards, boards.Items...)

	for len(boards.Bookmark) > 0 {
		boards, err = c.GetBoard(boards.Bookmark)
		if err != nil {
			return nil, err
		}

		resultBoards = append(resultBoards, boards.Items...)
	}

	return resultBoards, nil
}

func (c *Client) SearchPinsByaGivenSearchTerm(search string) (*models.PinData, error) {
	responseBytes, err := c.NewRequest("/search/partner/pins/" + search + "/")
	if err != nil {
		return nil, err
	}

	var pin models.PinData
	err = json.Unmarshal(responseBytes, &pin)
	if err != nil {
		return nil, err
	}

	return &pin, nil
}