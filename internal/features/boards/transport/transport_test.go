package boards_transport

import (
	"bytes"
	"context"
	"encoding/json"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockBoardsService struct {
	mock.Mock
}

func newMockBoardsService() *mockBoardsService {
	return &mockBoardsService{}
}

func (s *mockBoardsService) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {

	args := s.Called(ctx, board)

	var res domain.Board

	if args.Get(0) != nil {
		res = args.Get(0).(domain.Board)
	}

	err := args.Error(1)

	return res, err
}

func (s *mockBoardsService) DeleteBoard(ctx context.Context, boardId string, userId string) error {

	args := s.Called(ctx, boardId, userId)

	err := args.Error(0)

	return err
}

func (s *mockBoardsService) GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error) {

	args := s.Called(ctx, ownerId)

	var res []domain.Board

	if args.Get(0) != nil {
		res = args.Get(0).([]domain.Board)
	}

	err := args.Error(1)

	return res, err
}

func (s *mockBoardsService) GetSections(ctx context.Context, boardId string) (string, []domain.Section, error) {

	args := s.Called(ctx, boardId)

	var res []domain.Section

	if args.Get(1) != nil {
		res = args.Get(1).([]domain.Section)
	}

	err := args.Error(2)

	return args.String(0), res, err
}

type mockTokenManager struct {
	mock.Mock
}

func newMockTokenManager() *mockTokenManager {
	return &mockTokenManager{}
}

func (m *mockTokenManager) Generate(claims core_service_tokens.TokenClaims, tokenType core_service_tokens.TokenType, ttl time.Duration) (string, error) {

	args := m.Called(claims, tokenType, ttl)

	var res string

	if args.Get(0) != nil {
		res = args.Get(0).(string)
	}

	err := args.Error(1)

	return res, err
}

func (m *mockTokenManager) Parse(token string, tokenType core_service_tokens.TokenType) (*core_service_tokens.TokenClaims, error) {

	args := m.Called(token, tokenType)

	var res *core_service_tokens.TokenClaims

	if args.Get(0) != nil {
		res = args.Get(0).(*core_service_tokens.TokenClaims)
	}

	err := args.Error(1)

	return res, err
}

func TestBoardsHandler_CreateBoard(t *testing.T) {
	mockService := newMockBoardsService()
	mockTokenManager := newMockTokenManager()

	expectedBoard := domain.Board{ID: "123", Title: "Тестовая доска"}

	handler := NewBoardsHandler(mockService, mockTokenManager)

	mockService.On("CreateBoard", mock.Anything, mock.MatchedBy(func(b domain.Board) bool {
		return b.Title == "Тестовая доска"
	})).Return(expectedBoard, nil)

	requestBody, err := json.Marshal(map[string]string{"title": "Тестовая доска"})

	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/boards", bytes.NewBuffer(requestBody))

	testLogger := &core_logger.Logger{
		Logger: zap.NewNop(),
	}

	testClaims := &core_service_tokens.TokenClaims{
		UserID: "owner-1",
	}

	ctx := req.Context()
	ctx = core_logger.ToContext(ctx, testLogger)
	ctx = core_service_tokens.WithClaims(ctx, testClaims)
	req = req.WithContext(ctx)

	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.CreateBoard(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)

	var actualBoard domain.Board

	err = json.Unmarshal(rr.Body.Bytes(), &actualBoard)

	require.NoError(t, err)

	require.Equal(t, expectedBoard.ID, actualBoard.ID)
	require.Equal(t, expectedBoard.Title, actualBoard.Title)

	mockService.AssertExpectations(t)
	mockTokenManager.AssertExpectations(t)
}

func TestBoardsHandler_GetSections(t *testing.T) {
	mockService := newMockBoardsService()
	mockTokenManager := newMockTokenManager()

	boardID := "11111111-1111-1111-1111-111111111111"
	expectedSections := []domain.Section{
		{ID: "sec-1", Title: "Секция 1", Color: "red", BoardID: boardID},
	}

	handler := NewBoardsHandler(mockService, mockTokenManager)

	mockService.On("GetSections", mock.Anything, boardID).Return("board-title", expectedSections, nil)

	req, err := http.NewRequest(http.MethodGet, "/boards/"+boardID+"/sections", nil)
	require.NoError(t, err)
	req.SetPathValue("id", boardID)

	testLogger := &core_logger.Logger{
		Logger: zap.NewNop(),
	}
	req = req.WithContext(core_logger.ToContext(req.Context(), testLogger))

	rr := httptest.NewRecorder()

	handler.GetSections(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var actual GetSectionsResponse
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	require.NoError(t, err)
	require.Equal(t, "board-title", actual.BoardTitle)
	require.Len(t, actual.Sections, 1)
	require.Equal(t, "sec-1", actual.Sections[0].ID)

	mockService.AssertExpectations(t)
}
