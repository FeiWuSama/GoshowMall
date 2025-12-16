import axios from 'axios'
import { message } from 'ant-design-vue'
import Cookies from 'js-cookie'

// 创建 Axios 实例
const myAxios = axios.create({
  baseURL: '/api',
  timeout: 60000,
  withCredentials: true,
})

// 全局请求拦截器
myAxios.interceptors.request.use(
  function (config) {
    // 从localStorage获取captcha ticket
    const ticket = localStorage.getItem('captchaTicket')
    const expireStr = localStorage.getItem('captchaTicketExpire')
    // 检查ticket是否存在且未过期
    if (ticket) {
      if (expireStr) {
        const expireTime = parseInt(expireStr, 10)
        const currentTime = new Date().getTime()

        // 如果ticket未过期，添加到请求头
        if (currentTime < expireTime) {
          config.headers['Captcha-Ticket'] = ticket
        } else {
          // 如果过期，清除localStorage中的数据
          localStorage.removeItem('captchaTicket')
          localStorage.removeItem('captchaTicketExpire')
        }
      } else {
        // 如果没有过期时间，直接添加到请求头
        config.headers['Captcha-Ticket'] = ticket
      }
    }

    // 添加token
    if(Cookies.get('token')){
      config.headers['token'] = Cookies.get('token')
    }

    return config
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error)
  },
)

// 全局响应拦截器
myAxios.interceptors.response.use(
  function (response) {
    const { data } = response
    // 未登录
    if (data.code === 40100) {
      // 不是获取用户信息的请求，并且用户目前不是已经在用户登录页面，则跳转到登录页面
      if (
        !response.request.responseURL.includes('user/get/login') &&
        !window.location.pathname.includes('/user/login')
      ) {
        message.warning('请先登录')
        window.location.href = `/user/login?redirect=${window.location.href}`
      }
    }
    return response
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error)
  },
)

export default myAxios


