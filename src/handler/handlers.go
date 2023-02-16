package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmsic/deegle-translate-api/db/dbrepo"
)

// 全てのいいね数を取得
func AllLikes() gin.HandlerFunc {
	return func(c *gin.Context) {
		//実際の処理
		result := dbrepo.AllLikes()
		//json形式で画面に値返却
		c.JSON(http.StatusOK, gin.H{
			"likeCount": result,
		})
	}
}

type reqGTxt struct {
	Data `json:"data"`
}
type Data struct {
	Translations []Translations `json:"translations"`
}
type Translations struct {
	TranslatedText string `json:"translatedText"`
}

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
type TranslationResponse struct {
	Translations []Translation `json:"translations"`
}

func GoogleTranslate() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Printf("読み込み出来ませんでした: %v", err)
		}

		key := os.Getenv("NEXT_PUBLIC_GoogleTranslation_API_KEY")
		url := os.Getenv("NEXT_PUBLIC_GoogleTranslation_API_URL")

		text := c.Query("source")

		res, err := http.Get(url + "?key=" + key + "&q=" + text + "&source=JA&target=EN")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		// 取得したURLの内容を読み込む
		bytes, _ := io.ReadAll(res.Body)

		reqGTxt := reqGTxt{}
		if err := json.Unmarshal(bytes, &reqGTxt); err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"reqGTxt": reqGTxt,
		})

	}
}

func DeepLTranslate() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := godotenv.Load(".env")
		if err != nil {
			fmt.Printf("読み込み出来ませんでした: %v", err)
		}
		// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
		key := os.Getenv("NEXT_PUBLIC_DeepL_API_KEY")
		url := os.Getenv("NEXT_PUBLIC_DeepL_API_URL")

		text := c.Query("source")

		res, err := http.Get(url + "?" + "auth_key=" + key + "&text=" + text + "&source_lang=JA&target_lang=EN")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// レスポンスのJSONをパース
		var translationResponse TranslationResponse
		err = json.NewDecoder(res.Body).Decode(&translationResponse)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"reqDTxt": translationResponse,
		})
	}
}
