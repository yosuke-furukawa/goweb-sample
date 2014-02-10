package tweets

import (
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  "strconv"
  "net/http"
  "github.com/yosuke-furukawa/goweb-sample/models"
)
type TweetsController struct {
}

func (r *TweetsController) Read(id string, ctx context.Context) error {

  idInt, _ := strconv.Atoi(id)

  tweet, _ := tweets.FindById(idInt)
  if tweet.Id != 0 {
    return goweb.API.RespondWithData(ctx, tweet)
  } else {
    return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
  }
}

func (r *TweetsController) ReadMany(ctx context.Context) error {
  lastId := ctx.QueryValue("lastId")
  lastIdInt, _ := strconv.Atoi(lastId)
  tweetArr, _ := tweets.FindTweets(lastIdInt)
  return goweb.API.RespondWithData(ctx, tweetArr)
}

func (r *TweetsController) Delete(id string, ctx context.Context) error {

  idInt, _ := strconv.Atoi(id)
  tweets.DeleteById(idInt)
  return goweb.Respond.WithOK(ctx)
}

func (r *TweetsController) Create(ctx context.Context) error {
  data, dataErr := ctx.RequestData()

  if dataErr != nil {
    return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, dataErr.Error())
  }
  dataMap := data.(map[string]interface{})

  tweet := &tweets.Tweet{}
  tweet.Body = dataMap["body"].(string)
  tweet.Insert()
  return goweb.Respond.WithStatus(ctx, http.StatusCreated)
}
