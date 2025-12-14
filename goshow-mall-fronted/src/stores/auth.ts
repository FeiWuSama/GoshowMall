import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getUserInfo as apiGetUserInfo } from '@/api/user.ts'
import { message } from 'ant-design-vue'
import Cookies from 'js-cookie'

// 定义用户信息类型
interface UserInfo {
  id: number
  mobile?: string
  nickname?: string
  avatar?: string
  token: string
  sex: number
}

export const useAuthStore = defineStore('auth', () => {
  const isLoginModalVisible = ref(false)
  const userInfo = ref<UserInfo | null>(null)

  const openLoginModal = () => {
    isLoginModalVisible.value = true
  }

  const closeLoginModal = () => {
    isLoginModalVisible.value = false
  }

  // 登录成功，保存用户信息
  const loginSuccess = (userData: UserInfo) => {
    userInfo.value = userData
    // 保存token到cookies，用于页面刷新后自动登录
    Cookies.set('token', userData.token, { expires: 7 }) // 设置7天过期时间
  }

  // 登出
  const logout = () => {
    userInfo.value = null
    // 清除cookies中的token
    Cookies.remove('token')
  }

  // 页面加载时自动登录
  const autoLogin = async () => {
    // 从cookies获取token
    const token = Cookies.get('token')
    if (token) {
        try {
          const response = await apiGetUserInfo()
          if (response.code === 20000 && response.data) {
            // 更新用户信息
            userInfo.value = {
              avatar: response.data.avatar,
              id: response.data.id,
              mobile: response.data.mobile,
              nickname: response.data.nickname,
              sex: response.data.sex,
              token: response.data.token,
            }
          } else {
            message.error(response.msg || '获取用户信息失败')
            logout()
          }
        } catch (error: any) {
          message.error(error.message || '获取用户信息失败')
          logout()
        }
    }
  }

  return {
    isLoginModalVisible,
    userInfo,
    openLoginModal,
    closeLoginModal,
    loginSuccess,
    logout,
    autoLogin
  }
})
