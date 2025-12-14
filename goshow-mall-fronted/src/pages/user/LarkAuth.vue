<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { postUserLarkLogin } from '@/api/user'
import { useAuthStore } from '@/stores/auth'
import { message } from 'ant-design-vue'

const router = useRouter()
const authStore = useAuthStore()
const loading = ref(true)
const errorMsg = ref('')
const showSuccess = ref(false)

onMounted(() => {
  handleLarkAuth()
})

// 检查当前页面是否在iframe中
const isInIframe = () => {
  try {
    return window.self !== window.top
  } catch (e) {
    return true
  }
}

// 处理飞书授权回调
const handleLarkAuth = async () => {
  try {
    loading.value = true
    errorMsg.value = ''

    // 获取URL中的code参数
    const code = getQueryParam('code')
    const state = getQueryParam('state')

    if (!code) {
      throw new Error('未获取到飞书授权code')
    }

    // 发送飞书登录请求
    const response = await postUserLarkLogin({
      code,
      app_code: 1002,
      redirect_uri: `${window.location.origin}/user/lark/auth`,
    })

    if (response.code === 20000 && response.data) {
      // 登录成功，保存用户信息到Pinia
      authStore.loginSuccess({
        id: response.data.id,
        mobile: response.data.phone,
        nickname: response.data.nickname,
        avatar: response.data.avatar,
        token: response.data.token,
        sex: response.data.sex
      })

      // 延迟1.5秒后显示登录成功
      setTimeout(() => {
        showSuccess.value = true

        // 如果在iframe中，向父页面发送登录成功消息
        if (isInIframe()) {
          // 发送登录成功消息给父页面
          window.parent.postMessage(
            { type: 'lark-login-success', token: response.data.token },
            window.location.origin,
          )

          // 可以选择在iframe中关闭自身或显示成功信息
          // 这里我们保持显示成功信息，让父页面处理关闭
        } else {
          // 不在iframe中时，重定向到首页
          router.push('/')
        }
      }, 1500)
    } else {
      throw new Error(response.msg || '飞书登录失败')
    }
  } catch (error: any) {
    console.error('飞书授权登录失败:', error)
    errorMsg.value = error.message || '飞书登录失败，请重试'
    message.error(errorMsg.value)

    // 如果在iframe中，向父页面发送登录失败消息
    if (isInIframe()) {
      window.parent.postMessage(
        { type: 'lark-login-failed', error: errorMsg.value },
        window.location.origin,
      )
    } else {
      // 不在iframe中时，跳转到首页
      setTimeout(() => {
        router.push('/')
      }, 3000)
    }
  } finally {
    loading.value = false
  }
}

// 获取URL查询参数
const getQueryParam = (key: string): string | null => {
  const urlParams = new URLSearchParams(window.location.search)
  return urlParams.get(key)
}
</script>

<template>
  <div class="lark-auth-container">
    <div class="lark-auth-content">
      <h1 class="auth-title">飞书授权登录</h1>

      <div v-if="loading" class="loading-state">
        <div class="loading-spinner"></div>
        <p>正在验证飞书授权信息...</p>
      </div>

      <div v-else-if="errorMsg" class="error-state">
        <div class="error-icon">!</div>
        <p>{{ errorMsg }}</p>
        <p class="redirect-hint">即将返回首页...</p>
      </div>

      <div v-else-if="showSuccess" class="success-state">
        <div class="success-icon">✓</div>
        <p>登录成功，正在跳转...</p>
      </div>
      <div v-else class="processing-state">
        <div class="processing-spinner"></div>
        <p>正在处理登录信息...</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.lark-auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: white;
  padding: 20px;
}

.lark-auth-content {
  background: white;
  border-radius: 8px;
  padding: 40px;
  box-shadow: 0 2px 8px white;
  text-align: center;
  width: 100%;
  max-width: 400px;
  border: 1px solid white;
}

.auth-title {
  font-size: 20px;
  margin-bottom: 30px;
  color: #333;
  font-weight: 500;
}

.loading-spinner,
.processing-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid white;
  border-top: 3px solid #1890ff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 20px;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.error-icon {
  width: 40px;
  height: 40px;
  background: #f5f5f5;
  color: #ff4d4f;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  font-weight: bold;
  margin: 0 auto 20px;
  border: 2px solid #ff4d4f;
}

.success-icon {
  width: 40px;
  height: 40px;
  background: #f5f5f5;
  color: #52c41a;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  font-weight: bold;
  margin: 0 auto 20px;
  border: 2px solid #52c41a;
}

.loading-state p,
.error-state p,
.success-state p,
.processing-state p {
  font-size: 14px;
  color: #666;
  margin: 10px 0;
}

.redirect-hint {
  font-size: 12px;
  color: #999;
  margin-top: 20px;
}
</style>
