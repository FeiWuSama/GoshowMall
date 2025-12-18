// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** 获取滑块验证码 GET /api/user/captcha/slide */
export async function getUserCaptchaSlide(options?: { [key: string]: any }) {
  return request<API.ResultVoSlideCaptchaVo>('/api/user/captcha/slide', {
    method: 'GET',
    ...(options || {}),
  })
}

/** 验证滑块验证码 POST /api/user/captcha/slide/verify */
export async function postUserCaptchaSlideVerify(
  body: API.SlideCaptchaCheckDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultVoSlideCaptchaCheckVo>('/api/user/captcha/slide/verify', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 获取用户信息 GET /api/user/info */
export async function getUserInfo(options?: { [key: string]: any }) {
  return request<API.ResultVoUserVo>('/api/user/info', {
    method: 'GET',
    ...(options || {}),
  })
}

/** 飞书登录 POST /api/user/lark/login */
export async function postUserLarkLogin(
  body: API.UserLarkLoginDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultVoUserVo>('/api/user/lark/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 手机号登录 POST /api/user/mobile/login/password */
export async function postUserMobileLoginPassword(
  body: API.UserMobilePasswordLoginDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultVoUserVo>('/api/user/mobile/login/password', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 手机号短信验证码登录 POST /api/user/mobile/login/smsCode */
export async function postUserMobileLoginSmsCode(
  body: API.UserMobileSmsLoginDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultVoUserVo>('/api/user/mobile/login/smsCode', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 发送手机号短信验证码 POST /api/user/mobile/smsCode */
export async function postUserMobileSmsCode(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.postUserMobileSmsCodeParams,
  options?: { [key: string]: any }
) {
  return request<API.ResultAny>('/api/user/mobile/smsCode', {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  })
}

/** 注册 POST /api/user/register */
export async function postUserRegister(
  body: API.UserRegisterDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultAny>('/api/user/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}
