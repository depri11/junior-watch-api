package user_service_test

import (
	"context"
	"testing"

	"github.com/depri11/junior-watch-api/pkg/logger"
	go_proto "github.com/depri11/junior-watch-api/pkg/proto"
	"github.com/depri11/junior-watch-api/pkg/test"
	"github.com/depri11/junior-watch-api/user_service/internal/user/delivery"
	"github.com/depri11/junior-watch-api/user_service/internal/user/repository"
	"github.com/depri11/junior-watch-api/user_service/internal/user/service"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_DeliveryCreateUser(t *testing.T) {
	bt, err := test.NewBaseTest()
	assert.Nil(t, err)
	assert.NotNil(t, bt)

	repo := repository.NewUserRepository(&bt.Log, bt.Db)
	service := service.NewUserService(&bt.Log, repo)
	delivery := delivery.NewUserDelivery(service, bt.Log)

	type fields struct {
		log logger.Logger
		db  *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		user *go_proto.CreateUserRequest
	}

	f := fields{
		log: bt.Log,
		db:  bt.Db,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "should save a user",
			fields: f,
			args: args{
				ctx: context.Background(),
				user: &go_proto.CreateUserRequest{
					Username: "test",
					Email:    "test@test",
					Role:     go_proto.Role_USER,
					Name:     "test",
					Address:  "test",
					Phone:    "31351313",
				},
			},
			want:    "test-id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := delivery.Register(context.Background(), tt.args.user)
			assert.Nil(t, err)
			assert.NotEqual(t, "", res.UserID)
		})
	}
}
