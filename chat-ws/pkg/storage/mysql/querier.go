package mysql

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
)

type Querier interface {
	// user
	CreateUserDetail(ctx context.Context, params CreateUserDetailParam) error
	CreateUserCredential(ctx context.Context, params CreateUserCredentialParam) error
	CreateUserAccessInfo(ctx context.Context, params CreateUserAccessInfo) error
	GetUserDetailByUsername(ctx context.Context, username string) (account.UserDetail, error)
	GetUserDetailByUserId(ctx context.Context, userId string) (account.UserDetail, error)
	GetUserInfoByUserId(ctx context.Context, userId string) (chatting.UserDetail, error)
	GetUserDetailByEmail(ctx context.Context, email string) (account.UserDetail, error)
	GetUserCredentialByUsername(ctx context.Context, username string) (auth.User, error)
	GetUserCredentialByUserId(ctx context.Context, userId string) (auth.User, error)
	UpdateUser(ctx context.Context, params account.UpdateUserParam) error
	UpdateAccountStatus(ctx context.Context, userId, status string) error
	UpdateAccountPassword(ctx context.Context, userId, password string) error
	DeleteUser(ctx context.Context, userId string) error

	// verify_email
	CreateVerifyEmailToken(ctx context.Context, token account.VerifyEmailToken) error
	GetVerifyEmailToken(ctx context.Context, userId string) (account.VerifyEmailToken, error)
	DeleteVerifyEmailToken(ctx context.Context, tokenId string) error

	// reset_passwd
	CreateResetPasswordToken(ctx context.Context, params account.ResetPasswordToken) error
	GetResetPasswordToken(ctx context.Context, tokenId string) (account.ResetPasswordToken, error)
	GetResetPasswordTokenByUserId(ctx context.Context, userId string) (account.ResetPasswordToken, error)
	DeleteResetPasswordToken(ctx context.Context, tokenId string) error

	// workspace
	CreateWorkspace(ctx context.Context, w adding.Workspace) error
	AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (listing.Workspace, error)
	GetWorkspacesByUserId(ctx context.Context, userId string) ([]listing.Workspace, error)
	UpdateWorkspace(ctx context.Context, workspace editing.Workspace) error
	DeleteWorkspace(ctx context.Context, id string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error

	// channel
	CreateChannel(ctx context.Context, channel adding.Channel) error
	AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error
	GetChannelByChannelId(ctx context.Context, channelId string) (listing.Channel, error)
	GetChannelsByUserId(ctx context.Context, userId string) ([]listing.Channel, error)
	GetChannelListByWorkspaceId(ctx context.Context, workspaceId string) ([]listing.Channel, error)
	// GetUserListByChannelId(ctx context.Context, channelId string) ([]chatting.UserData, error)
	GetDefaultChannel(ctx context.Context, workspaceId string) (listing.Channel, error)
	GetDefaultChannelByWorkspaceId(ctx context.Context, workspaceId string) (adding.Channel, error)
	UpdateChannel(ctx context.Context, channel editing.Channel) error
	DeleteChannel(ctx context.Context, channelId string) error
	RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error

	// message
	CreateMessage(ctx context.Context, message *chatting.Message) error
	GetChannelMessages(ctx context.Context, channelId string, limit, offset uint64) ([]chatting.Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]chatting.Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]chatting.Message, error)
	UpdateChannelAccessTime(ctx context.Context, channelId string, userId string) error

	// websocket
	GetUserListByChannelId(ctx context.Context, channelId string) ([]websocket.UserData, error)
	GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]websocket.UserData, error)

	// board
	CreateBoard(ctx context.Context, board board.TaskBoard) error
	GetTaskBoardByWorkspaceId(ctx context.Context, workspaceId string) (board.TaskBoard, error)
	UpdateTaskBoard(ctx context.Context, board board.TaskBoard) error

	CreateTaskList(ctx context.Context, taskList board.TaskList) error
	GetTaskListsByBoardId(ctx context.Context, boardId string) ([]board.TaskList, error)
	UpdateTaskList(ctx context.Context, taskList board.TaskList) error
	UpdateTaskListOrder(ctx context.Context, listId string, orderIndex int) error
	DeleteTaskList(ctx context.Context, listId string) error

	CreateCard(ctx context.Context, card board.Card) error
	GetCardListByListId(ctx context.Context, listId string) ([]board.Card, error)
	UpdateCard(ctx context.Context, params board.UpdateCardParams) error
	UpdateCardTitle(ctx context.Context, params board.UpdateCardTitleParams) error
	UpdateCardOrder(ctx context.Context, cardId string, orderIndex int, listId string, updateListId bool) error
	DeleteCard(ctx context.Context, cardId string) error
	AddUserToCard(ctx context.Context, cardId, userId string) error
	DeleteUserFromCard(ctx context.Context, cardId, userId string) error

	CreateCardComment(ctx context.Context, comment *board.CardComment) error
	GetCardComments(ctx context.Context, cardId string) ([]board.CardComment, error)
}
