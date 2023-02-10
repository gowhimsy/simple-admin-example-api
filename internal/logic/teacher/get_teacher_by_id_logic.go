package teacher

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetTeacherByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetTeacherByIdLogic {
	return &GetTeacherByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetTeacherByIdLogic) GetTeacherById(req *types.UUIDReq) (resp *types.TeacherInfoResp, err error) {
	data, err := l.svcCtx.SchoolRpc.GetTeacherById(l.ctx, &example.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TeacherInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.TeacherInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:          data.Name,
			Age:           data.Age,
			AgeInt32:      data.AgeInt32,
			AgeInt64:      data.AgeInt64,
			AgeUint:       data.AgeUint,
			AgeUint32:     data.AgeUint32,
			AgeUint64:     data.AgeUint64,
			WeightFloat:   data.WeightFloat,
			WeightFloat32: data.WeightFloat32,
			ClassId:       data.ClassId,
			EnrollAt:      data.EnrollAt,
			StatusBool:    data.StatusBool,
		},
	}, nil
}
