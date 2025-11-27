package captcha

import (
	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
)

func NewCaptcha() slide.Captcha {
	builder := slide.NewBuilder(slide.WithGenGraphNumber(1))
	background, err := imagesv2.GetImages()
	if err != nil {
		panic(err)
	}
	graphImages, err := tiles.GetTiles()
	if err != nil {
		panic(err)
	}
	var newImages = make([]*slide.GraphImage, 0, len(graphImages))
	for _, image := range graphImages {
		newImages = append(newImages, &slide.GraphImage{
			MaskImage:    image.MaskImage,
			OverlayImage: image.OverlayImage,
			ShadowImage:  image.ShadowImage,
		})
	}

	builder.SetResources(
		slide.WithBackgrounds(background),
		slide.WithGraphImages(newImages),
	)
	return builder.Make()
}
