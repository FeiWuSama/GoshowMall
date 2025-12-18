// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** 获取滑块验证码 GET /api/admin/captcha/slide */
export async function getAdminCaptchaSlide(options?: { [key: string]: any }) {
  return request<API.ResultVoSlideCaptchaVo>('/api/admin/captcha/slide', {
    method: 'GET',
    ...(options || {}),
  })
}

/** 验证滑块验证码 POST /api/admin/captcha/slide/verify */
export async function postAdminCaptchaSlideVerify(
  body: API.SlideCaptchaCheckDto,
  options?: { [key: string]: any }
) {
  return request<API.ResultVoSlideCaptchaCheckVo>('/api/admin/captcha/slide/verify', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 创建用户 POST /api/admin/create */
export async function postAdminCreate(body: API.AddAdminDto, options?: { [key: string]: any }) {
  return request<any>('/api/admin/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 获取管理员信息 GET /api/admin/info */
export async function getAdminInfo(options?: { [key: string]: any }) {
  return request<API.ResultVoAdminVO>('/api/admin/info', {
    method: 'GET',
    ...(options || {}),
  })
}

/** 管理员登录 POST /api/admin/login */
export async function postAdminLogin(body: API.AdminLoginDto, options?: { [key: string]: any }) {
  return request<API.ResultVoAdminVO>('/api/admin/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** 改变用户状态 POST /api/admin/status/${param0}/${param1} */
export async function postAdminStatusIdStatus(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.postAdminStatusIdStatusParams,
  options?: { [key: string]: any }
) {
  const { id: param0, status: param1, ...queryParams } = params
  return request<any>(`/api/admin/status/${param0}/${param1}`, {
    method: 'POST',
    params: { ...queryParams },
    ...(options || {}),
  })
}

/** 更新用户 POST /api/admin/update */
export async function postAdminUpdate(body: API.UpdateAdminDto, options?: { [key: string]: any }) {
  return request<any>('/api/admin/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}
