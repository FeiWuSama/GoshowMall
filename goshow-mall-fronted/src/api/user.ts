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
