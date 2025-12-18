declare namespace API {
  type AddAdminDto = {
    mobile?: string
    nickname?: string
    sex?: number
    username?: string
  }

  type AdminLoginDto = {
    mobile?: string
    password?: string
  }

  type AdminVO = {
    id?: number
    name?: string
    nickname?: string
    sex?: number
    token?: string
  }

  type postAdminStatusIdStatusParams = {
    /** 用户id */
    id: number
    /** 用户状态 */
    status: number
  }

  type postUserMobileSmsCodeParams = {
    /** ticket */
    ticket: string
    /** mobile */
    mobile: string
    /** 场景 */
    scene: string
  }

  type ResultAny = {
    code?: number
    data?: any
    msg?: string
  }

  type ResultVoAdminVO = {
    code?: number
    data?: AdminVO
    msg?: string
  }

  type ResultVoSlideCaptchaCheckVo = {
    code?: number
    data?: SlideCaptchaCheckVo
    msg?: string
  }

  type ResultVoSlideCaptchaVo = {
    code?: number
    data?: SlideCaptchaVo
    msg?: string
  }

  type ResultVoUserVo = {
    code?: number
    data?: UserVo
    msg?: string
  }

  type SlideCaptchaCheckDto = {
    key?: string
    slideX?: number
    slideY?: number
  }

  type SlideCaptchaCheckVo = {
    expire?: number
    ticket?: string
  }

  type SlideCaptchaVo = {
    ImageBase64?: string
    TitleHeight?: number
    TitleImageBase64?: string
    TitleWidth?: number
    TitleX?: number
    TitleY?: number
    key?: string
  }

  type UpdateAdminDto = {
    id?: number
    mobile?: string
    nickname?: string
    sex?: number
    username?: string
  }

  type UserLarkLoginDto = {
    app_code?: number
    code?: string
    redirect_uri?: string
  }

  type UserMobilePasswordLoginDto = {
    mobile?: string
    password?: string
    ticket?: string
  }

  type UserMobileSmsLoginDto = {
    mobile?: string
    scene?: string
    verify_code?: string
  }

  type UserRegisterDto = {
    mobile?: string
    nickname?: string
    password?: string
    scene?: string
    sex?: number
    verify_code?: string
  }

  type UserVo = {
    avatar?: string
    id?: number
    nickname?: string
    sex?: number
    token?: string
  }
}
