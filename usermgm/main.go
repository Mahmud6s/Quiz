package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	answerpb "quiz/gunk/v1/answer"
	categorypb "quiz/gunk/v1/category"
	optionpb "quiz/gunk/v1/option"
	questionpb "quiz/gunk/v1/question"
	quizpb "quiz/gunk/v1/quiz"
	qzquestionpb "quiz/gunk/v1/quiz_question"
	userpb "quiz/gunk/v1/user"
	userquizpb "quiz/gunk/v1/user_quiz"
	"strings"

	cc "quiz/usermgm/core/category"
	cu "quiz/usermgm/core/user"

	as "quiz/usermgm/core/answer"
	"quiz/usermgm/service/answer"
	"quiz/usermgm/service/category"
	"quiz/usermgm/service/option"
	"quiz/usermgm/service/quizquestion"
	"quiz/usermgm/service/userquiz"

	uq "quiz/usermgm/core/userquiz"

	qzq "quiz/usermgm/core/quizquestion"

	"quiz/usermgm/service/quiz"

	qz "quiz/usermgm/core/quiz"

	oo "quiz/usermgm/core/option"
	"quiz/usermgm/service/question"

	qq "quiz/usermgm/core/question"
	"quiz/usermgm/service/user"
	"quiz/usermgm/storage/postgres"

	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:embed migrations
var migrationFiles embed.FS

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	port := config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen port: %v", err)
	}

	postgreStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}

	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err := goose.Up(postgreStorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	//..for user
	userCore := cu.NewCoreUser(postgreStorage)
	userSvc := user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer, userSvc)
	//..for Category
	categoryCore := cc.NewCoreCategory(postgreStorage)
	categorySvc := category.NewCategorySvc(categoryCore)
	categorypb.RegisterCategoryServiceServer(grpcServer, categorySvc)
	//for questios
	questionCore := qq.NewCoreQuestion(postgreStorage)
	questionSvc := question.NewQuestionSvc(questionCore)
	questionpb.RegisterQuestionServiceServer(grpcServer, questionSvc)
	//for Options
	optionCore := oo.NewCoreOption(postgreStorage)
	optionSvc := option.NewOptionSvc(optionCore)
	optionpb.RegisterOptionServiceServer(grpcServer, optionSvc)
	//for Quiz
	quizCore := qz.NewCoreQuiz(postgreStorage)
	quizSvc := quiz.NewQuizSvc(quizCore)
	quizpb.RegisterQuizServiceServer(grpcServer, quizSvc)
	//for Quiz_questionn
	qzquestionCore := qzq.NewCoreQuestion(postgreStorage)
	qzquestionSvc := quizquestion.NewQuizQuestionSvc(qzquestionCore)
	qzquestionpb.RegisterQuizQuestionServiceServer(grpcServer, qzquestionSvc)
	//..for User_Quiz
	userquizCore := uq.NewCoreUserQuiz(postgreStorage)
	userquizSvc := userquiz.NewUserQuizSvc(userquizCore)
	userquizpb.RegisterUserQuizServiceServer(grpcServer, userquizSvc)
	//..for Answer
	answerCore := as.NewCoreAnswer(postgreStorage)
	answerSvc := answer.NewAnswerSvc(answerCore)
	answerpb.RegisterAnswerServiceServer(grpcServer, answerSvc)
	// start reflection server

	reflection.Register(grpcServer)
	fmt.Println("usermgm server running on: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}
