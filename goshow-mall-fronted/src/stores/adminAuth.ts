import { defineStore } from 'pinia'
import { ref } from 'vue'
import {  getAdminInfo } from '@/api/admin.ts'
import { message } from 'ant-design-vue'
import Cookies from 'js-cookie'

// 定义管理员信息类型
interface AdminInfo {
  id: number
  nickname?: string
  token: string
}

export const useAdminAuthStore = defineStore('adminAuth', () => {
  const adminInfo = ref<AdminInfo | null>(null)

  // 管理员登录成功，保存管理员信息
  const loginSuccess = (adminData: AdminInfo) => {
    adminInfo.value = adminData
    // 保存token到cookies，用于页面刷新后自动登录
    Cookies.set('admin-token', adminData.token, { expires: 7 }) // 设置7天过期时间
  }

  // 管理员登出
  const logout = () => {
    adminInfo.value = null
    // 清除cookies中的token
    Cookies.remove('admin-token')
  }

  // 页面加载时自动登录管理员
  const autoLogin = async () => {
    // 从cookies获取token
    const token = Cookies.get('admin-token')
    if (token) {
        try {
          const response = await getAdminInfo()
          if (response.data.code === 20000 && response.data.data) {
            // 更新管理员信息
            adminInfo.value = {
              id: response.data.data.id || 0,
              nickname: response.data.data.nickname,
              token: response.data.data.token || '',
            }
          } else {
            message.error(response.data.msg || '获取管理员信息失败')
            logout()
          }
        } catch (error: any) {
          message.error(error.message || '获取管理员信息失败')
          logout()
        }
    }
  }

  return {
    adminInfo,
    loginSuccess,
    logout,
    autoLogin
  }
})
