<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router'
import { useAuthStore } from './stores/auth'
import Layout from './layout/Layout.vue'
import LoginModal from './components/LoginModal.vue'

const authStore = useAuthStore()
const route = useRoute()

const handleLoginSuccess = (token: string) => {
  console.log('用户已登录，token:', token)
  authStore.closeLoginModal()
}

// 需要排除布局的路由路径
const noLayoutRoutes = ['/user/lark/auth']
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
