<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router'
import { useAuthStore } from './stores/auth'
import Layout from './layout/Layout.vue'
import LoginModal from './components/LoginModal.vue'
import { onMounted } from 'vue'

const authStore = useAuthStore()
const route = useRoute()

const handleLoginSuccess = (userVo: API.UserVo) => {
  authStore.loginSuccess({
    id: userVo.id || 0,
    nickname: userVo.nickname,
    avatar: userVo.avatar,
    token: userVo.token || '',
    sex: userVo.sex || 0,
  })
  authStore.closeLoginModal()
}

// 需要排除布局的路由路径
const noLayoutRoutes = ['/user/lark/auth']

// 页面挂载时自动登录
onMounted(() => {
  // 从cookies中获取token并自动登录
  authStore.autoLogin()
})
</script>

<template>
  <!-- 根据当前路由决定是否显示布局 -->
  <template v-if="!noLayoutRoutes.includes(route.path)">
    <Layout>
      <RouterView />
    </Layout>
  </template>
  <template v-else>
    <RouterView />
  </template>
  <!-- 全局登录弹窗 -->
  <LoginModal
    :visible="authStore.isLoginModalVisible"
    @update:visible="authStore.closeLoginModal"
    @login-success="handleLoginSuccess"
  />
</template>

<style>
#app {
  width: 100%;
  height: 100%;
}
</style>
