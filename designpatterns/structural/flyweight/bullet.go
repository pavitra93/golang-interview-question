package flyweight

type bullet struct {
	radius     float64
	length     float64
	image      string
	bulletType BulletType
	metalType  MetalType
}

type BulletBuilder struct {
	bullet *bullet
}

func GetInstance() *BulletBuilder {
	return &BulletBuilder{bullet: &bullet{}}
}

func (builder *BulletBuilder) SetRadius(radius float64) *BulletBuilder {
	builder.bullet.radius = radius
	return builder
}

func (builder *BulletBuilder) SetLength(length float64) *BulletBuilder {
	builder.bullet.length = length
	return builder
}

func (builder *BulletBuilder) SetBulletType(bulletType BulletType) *BulletBuilder {
	builder.bullet.bulletType = bulletType
	return builder
}

func (builder *BulletBuilder) SetMetal(metalType MetalType) *BulletBuilder {
	builder.bullet.metalType = metalType
	return builder
}

func (builder *BulletBuilder) SetImage(image string) *BulletBuilder {
	builder.bullet.image = image
	return builder
}

func (builder *BulletBuilder) Build() *bullet {
	return &bullet{
		radius:     builder.bullet.radius,
		length:     builder.bullet.length,
		image:      builder.bullet.image,
		bulletType: builder.bullet.bulletType,
		metalType:  builder.bullet.metalType,
	}
}
