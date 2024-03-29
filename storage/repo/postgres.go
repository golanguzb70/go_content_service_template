package repo

import (
	"context"

	pb "github.com/golanguzb70/go_content_service/genproto/content_service"
)

type PositionI interface {
	Create(ctx context.Context, req *pb.Position) (*pb.Position, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Position, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Positions, error)
	Update(ctx context.Context, req *pb.Position) (*pb.Position, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type StaffI interface {
	Create(ctx context.Context, req *pb.Staff) (*pb.Staff, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Staff, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Staffs, error)
	Update(ctx context.Context, req *pb.Staff) (*pb.Staff, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type TagI interface {
	Create(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Tag, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Tags, error)
	Update(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type GenreI interface {
	Create(ctx context.Context, req *pb.Genre) (*pb.Genre, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Genre, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Genres, error)
	Update(ctx context.Context, req *pb.Genre) (*pb.Genre, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type CategoryI interface {
	Create(ctx context.Context, req *pb.Category) (*pb.Category, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Category, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Categories, error)
	Update(ctx context.Context, req *pb.Category) (*pb.Category, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type CountryI interface {
	Create(ctx context.Context, req *pb.Country) (*pb.Country, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Country, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Countries, error)
	Update(ctx context.Context, req *pb.Country) (*pb.Country, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type ContentProviderI interface {
	Create(ctx context.Context, req *pb.ContentProvider) (*pb.ContentProvider, error)
	Get(ctx context.Context, req *pb.Id) (*pb.ContentProvider, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.ContentProviders, error)
	Update(ctx context.Context, req *pb.ContentProvider) (*pb.ContentProvider, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}
