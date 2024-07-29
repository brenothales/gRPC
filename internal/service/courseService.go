package service

import (
	"context"
	"io"

	"github.com/brenothales/gRPC/internal/entity"
	"github.com/brenothales/gRPC/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB entity.Course
}

func NewCourseService(CourseDB entity.Course) *CourseService {
	return &CourseService{
		CourseDB: CourseDB,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.Course, error) {
	Course, err := c.CourseDB.Create(in.Name, in.Description, in.CategoryID)
	if err != nil {
		return nil, err
	}

	CourseResponse := &pb.Course{
		Id:          Course.ID,
		Name:        Course.Name,
		Description: Course.Description,
	}

	return CourseResponse, nil
}

func (c *CourseService) ListCourses(ctx context.Context, in *pb.CourseBlank) (*pb.CourseList, error) {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}

	var coursesResponse []*pb.Course

	for _, Course := range courses {
		CourseResponse := &pb.Course{
			Id:          Course.ID,
			Name:        Course.Name,
			Description: Course.Description,
		}

		coursesResponse = append(coursesResponse, CourseResponse)
	}

	return &pb.CourseList{Courses: coursesResponse}, nil
}

func (c *CourseService) ListCoursesStream(in *pb.CourseBlank, stream pb.CourseService_ListCoursesStreamServer) error {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return err
	}

	for _, course := range courses {
		courseResponse := &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		}

		if err := stream.Send(courseResponse); err != nil {
			return err
		}
	}

	return nil
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.CourseGetRequest) (*pb.Course, error) {
	Course, err := c.CourseDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	CourseResponse := &pb.Course{
		Id:          Course.ID,
		Name:        Course.Name,
		Description: Course.Description,
	}

	return CourseResponse, nil
}

func (c *CourseService) CreateCourseStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CourseList{}

	for {
		Course, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}
		if err != nil {
			return err
		}

		CourseResult, err := c.CourseDB.Create(Course.Name, Course.Description, Course.CategoryID)
		if err != nil {
			return err
		}

		courses.Courses = append(courses.Courses, &pb.Course{
			Id:          CourseResult.ID,
			Name:        CourseResult.Name,
			Description: CourseResult.Description,
		})
	}
}

func (c *CourseService) CreateCourseStreamBidirectional(stream pb.CourseService_CreateCourseStreamBidirectionalServer) error {
	for {
		Course, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		CourseResult, err := c.CourseDB.Create(Course.Name, Course.Description, Course.CategoryID)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Course{
			Id:          CourseResult.ID,
			Name:        CourseResult.Name,
			Description: CourseResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
