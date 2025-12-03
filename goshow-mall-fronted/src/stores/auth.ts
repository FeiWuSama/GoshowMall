import { defineStore } from 'pinia'
import { ref } from 'vue'

// 定义用户信息类型
interface UserInfo {
  id: number
  phone?: string
  nickname?: string
  avatar?: string
  token: string
  [key: string]: any
}

export const useAuthStore = defineStore('auth', () => {
  const isLoginModalVisible = ref(false)
  const userInfo = ref<UserInfo | null>(null)
  const isLoggedIn = ref(false)

  const openLoginModal = () => {
    isLoginModalVisible.value = true
  }

  const closeLoginModal = () => {
    isLoginModalVisible.value = false
  }

  // 登录成功，保存用户信息
  const loginSuccess = (userData: UserInfo) => {
    userInfo.value = userData
    isLoggedIn.value = true
    // 保存到localStorage，实现持久化存储
    localStorage.setItem('userInfo', JSON.stringify(userData))
    localStorage.setItem('isLoggedIn', 'true')
  }

  // 登出
  const logout = () => {
    userInfo.value = null
    isLoggedIn.value = false
    // 清除localStorage中的数据
    localStorage.removeItem('userInfo')
    localStorage.removeItem('isLoggedIn')
  }

  // 初始化时从localStorage恢复用户信息
  const initUserInfo = () => {
    const savedUserInfo = localStorage.getItem('userInfo')
    const savedIsLoggedIn = localStorage.getItem('isLoggedIn')

    if (savedUserInfo && savedIsLoggedIn === 'true') {
      try {
        userInfo.value = JSON.parse(savedUserInfo)
        isLoggedIn.value = true
      } catch (error) {
        // 如果解析失败，清除数据
        logout()
      }
    }
  }

  return {
    isLoginModalVisible,
    userInfo,
    isLoggedIn,
    openLoginModal,
    closeLoginModal,
    loginSuccess,
    logout,
    initUserInfo
  }
})
