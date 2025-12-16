<script setup lang="ts">
import { ref, onMounted, watch, nextTick, onUnmounted } from 'vue'
import { Form, Input, Button, Tabs, Space, message, Modal } from 'ant-design-vue'
import { LockOutlined, PhoneOutlined } from '@ant-design/icons-vue'
import {
  getUserCaptchaSlide,
  postUserCaptchaSlideVerify,
  postUserMobileLoginPassword,
  postUserMobileSmsCode,
  postUserMobileLoginSmsCode,
} from '@/api/user.ts'
import { useAuthStore } from '@/stores/auth.ts'
import { CODE_SCENE } from '@/constant/constant.ts'

interface Props {
  visible: boolean
}

const props = defineProps<Props>()

// 获取auth store
const authStore = useAuthStore()

const emit = defineEmits<{
  'update:visible': [value: boolean]
  'login-success': [userInfo: Object]
}>()

// 手机号密码登录
const accountForm = ref({
  mobile: '',
  password: '',
})

// 是否正在进行滑块验证码验证
const isCaptchaRequired = ref(false)

// 验证码登录
const codeForm = ref({
  phone: '',
  code: '',
})

const isLoading = ref(false)
const activeTab = ref('account')
const countDown = ref(0)
// 验证码计时器
let timer: number | null = null
// 当前验证码请求的场景
const captchaScene = ref('')

// 验证码相关状态
const captchaModal = ref(false)
const captchaData = ref<any>(null)
const captchaKey = ref('')
const captchaSliderPosition = ref(0)
const isDragging = ref(false)
const captchaVerifying = ref(false)
const sliderContainer = ref<HTMLElement | null>(null)

// 关闭弹窗
const handleClose = () => {
  emit('update:visible', false)
  resetForms()
}

// 重置表单
const resetForms = () => {
  accountForm.value = { mobile: '', password: '' }
  codeForm.value = { phone: '', code: '' }
  activeTab.value = 'account'
  countDown.value = 0
}

// 检查是否为手机号
const isPhoneNumber = (value: string): boolean => {
  const phoneRegex = /^1[3-9]\d{9}$/
  return phoneRegex.test(value)
}

// 手机号密码登录
const handleAccountLogin = async () => {
  if (!accountForm.value.mobile || !accountForm.value.password) {
    message.error('请填写手机号和密码')
    return
  }

  // 检查是否为手机号登录
  if (!isPhoneNumber(accountForm.value.mobile)) {
    message.error('请输入正确的手机号')
    return
  }

  // 如果是手机号，先获取滑块验证码
  try {
    const result = await getUserCaptchaSlide()
    if (result.data.code === 20000 && result.data.data) {
      captchaData.value = result.data.data
      captchaKey.value = result.data.data.key || ''
      captchaSliderPosition.value = 0 // 初始位置设置为0
      captchaModal.value = true // 打开滑块验证码弹窗
      isCaptchaRequired.value = true // 标记需要验证码验证
    } else {
      message.error('获取验证码失败')
    }
  } catch (error) {
    message.error('获取验证码失败')
  }
}

// 执行登录请求
const performLogin = async () => {
  isLoading.value = true
  try {
    // 手机号登录
    const response = await postUserMobileLoginPassword({
      mobile: accountForm.value.mobile,
      password: accountForm.value.password,
    })
    if (response.data.code === 20000 && response.data.data) {
      message.success('登录成功')

      // 保存用户信息到pinia store
      authStore.loginSuccess({
        id: response.data.data.id || 0,
        nickname: response.data.data.nickname,
        avatar: response.data.data.avatar,
        token: response.data.data.token || '',
        sex: response.data.data.sex || 0,
      })

      // 触发登录成功事件
      emit('login-success', response.data.data)
      handleClose()
    } else {
      message.error(response.data.msg || '登录失败')
    }
  } catch (error: any) {
    message.error(error.message || '登录失败')
  } finally {
    isLoading.value = false
  }
}

// 获取验证码
const handleGetCode = async () => {
  if (!codeForm.value.phone) {
    message.error('请输入手机号')
    return
  }

  // 检查是否为手机号
  if (!isPhoneNumber(codeForm.value.phone)) {
    message.error('请输入正确的手机号')
    return
  }

  try {
    // 设置当前场景为获取验证码
    captchaScene.value = 'getCode'
    // 先获取滑块验证码
    let result = await getUserCaptchaSlide()
    if (result.data.code === 20000 && result.data.data) {
      captchaData.value = result.data
      captchaKey.value = result.data.data.key || ''
      captchaSliderPosition.value = 0 // 初始位置设置为0
      captchaModal.value = true // 打开滑块验证码弹窗
    } else {
      message.error('获取验证码失败')
    }
  } catch (error) {
    message.error('获取验证码失败')
  }
}

// 验证码登录
const handleCodeLogin = async () => {
  if (!codeForm.value.phone || !codeForm.value.code) {
    message.error('请填写手机号和验证码')
    return
  }

  // 检查是否为手机号
  if (!isPhoneNumber(codeForm.value.phone)) {
    message.error('请输入正确的手机号')
    return
  }

  isLoading.value = true
  try {
    const response = await postUserMobileLoginSmsCode({
      mobile: codeForm.value.phone,
      verify_code: codeForm.value.code,
      scene: CODE_SCENE.LOGIN,
    })
    if (response.data.code === 20000 && response.data.data) {
      message.success('登录成功')
      console.log('登录响应:', response)

      // 保存用户信息到pinia store
      authStore.loginSuccess({
        id: response.data.data.id || 0,
        nickname: response.data.data.nickname,
        avatar: response.data.data.avatar,
        token: response.data.data.token || '',
        sex: response.data.data.sex || 0,
      })

      emit('login-success', response.data.data)
      message.success('登录成功')
      handleClose()
    } else {
      message.error(response.data.msg || '登录失败')
    }
  } catch (error: any) {
    message.error(error.message || '登录失败')
  } finally {
    isLoading.value = false
  }
}

// 验证滑块
const handleVerifyCaptcha = async () => {
  captchaVerifying.value = true
  try {
    // 计算拼图的实际位置，即滑块位置加上拼图初始位置
    const puzzleX = captchaSliderPosition.value

    // 发送验证请求
    const response = await postUserCaptchaSlideVerify({
      key: captchaKey.value,
      slideX: Math.round(puzzleX),
      slideY: Math.round(captchaData.value.TitleY), // 使用接口返回的Y坐标
    })
    if (response.data.code === 20000 && response.data.data) {
      // 保存ticket凭证到localStorage
      const { ticket, expire } = response.data.data
      if (ticket) {
        localStorage.setItem('captchaTicket', ticket)
        // 如果有过期时间，设置过期时间
        if (expire) {
          const expireTime = new Date().getTime() + expire * 60 * 1000
          localStorage.setItem('captchaTicketExpire', expireTime.toString())
        }
      }

      // 检查是否是为了登录而进行的验证
      if (isCaptchaRequired.value) {
        captchaModal.value = false
        // 验证码验证成功后，发送登录请求
        await performLogin()
        // 重置验证码验证标记
        isCaptchaRequired.value = false
      } else if (captchaScene.value === 'getCode') {
        // 如果是为了获取验证码而进行的验证
        captchaModal.value = false
        await sendSmsCode()
        // 重置场景
        captchaScene.value = ''
      } else {
        captchaModal.value = false
      }
    } else {
      message.error('验证失败')
      // 重置滑块位置
      captchaSliderPosition.value = 0
      // 重新获取滑块验证码
      await fetchNewCaptcha()
    }
  } catch (error) {
    message.error('验证失败')
    // 重置滑块位置
    captchaSliderPosition.value = 0
    // 重新获取滑块验证码
    await fetchNewCaptcha()
  } finally {
    captchaVerifying.value = false
  }
}

// 发送短信验证码
const sendSmsCode = async () => {
  try {
    // 调用发送验证码接口
    const response = await postUserMobileSmsCode({
      ticket: localStorage.getItem('captchaTicket') || '',
      mobile: codeForm.value.phone,
      scene: CODE_SCENE.LOGIN,
    })
    if (response.data.code === 20000) {
      message.success('验证码发送成功')
      // 开始倒计时
      startCountDown()
    } else {
      message.error(response.data.msg || '验证码发送失败')
    }
  } catch (error: any) {
    message.error(error.message || '验证码发送失败')
  }
}

// 开始倒计时
const startCountDown = () => {
  countDown.value = 60
  // 清除之前的计时器
  if (timer) {
    clearInterval(timer)
  }
  // 设置新的计时器
  timer = setInterval(() => {
    countDown.value--
    if (countDown.value <= 0) {
      // 倒计时结束，清除计时器
      if (timer) {
        clearInterval(timer)
        timer = null
      }
    }
  }, 1000) as unknown as number
}

// 重新获取滑块验证码
const fetchNewCaptcha = async () => {
  try {
    const result = await getUserCaptchaSlide()
    if (result.data.code === 20000 && result.data.data) {
      captchaData.value = result.data.data
      captchaKey.value = result.data.data.key || ''
      captchaSliderPosition.value = 0 // 初始位置设置为0
    } else {
      message.error('获取验证码失败，请稍后重试')
      captchaModal.value = false
    }
  } catch (error) {
    message.error('获取验证码失败，请稍后重试')
    captchaModal.value = false
  }
}

// 关闭slider验证码弹窗
const handleCloseCaptchaModal = () => {
  captchaModal.value = false
  captchaSliderPosition.value = 0
}

// 处理滑块按下事件
const handleSliderMouseDown = (event: MouseEvent) => {
  event.preventDefault()
  isDragging.value = true

  // 在开始拖动时缓存容器信息，避免拖动过程中频繁计算
  cacheContainerInfo()

  // 重置lastPosition
  lastPosition = captchaSliderPosition.value

  // 添加全局鼠标事件监听
  document.addEventListener('mousemove', handleGlobalMouseMove)
  document.addEventListener('mouseup', handleGlobalMouseUp)
}

// 用于requestAnimationFrame的位置更新函数
let animationFrameId: number
let lastPosition: number = -1
let containerInfo: { width: number; left: number } | null = null

// 缓存容器信息
const cacheContainerInfo = () => {
  if (!sliderContainer.value) return
  const rect = sliderContainer.value.getBoundingClientRect()
  containerInfo = {
    width: rect.width,
    left: rect.left,
  }
}

// 全局鼠标移动事件处理
const handleGlobalMouseMove = (event: MouseEvent) => {
  if (!isDragging.value) return

  event.preventDefault()

  // 使用requestAnimationFrame优化性能，避免频繁更新
  if (animationFrameId) return

  animationFrameId = requestAnimationFrame(() => {
    // 只在缓存为空时重新获取容器信息
    if (!containerInfo) {
      cacheContainerInfo()
    }

    if (!containerInfo || !captchaData.value) {
      animationFrameId = 0
      return
    }

    // 使用缓存的容器信息，减少getBoundingClientRect调用
    const titleWidth = captchaData.value.TitleWidth * 0.5

    // 计算新位置，确保滑块不会超出容器边界
    let newPosition = event.clientX - containerInfo.left - titleWidth
    newPosition = Math.max(0, Math.min(newPosition, containerInfo.width - 50))

    // 只有位置发生变化超过1px时才更新，避免极小的波动引起渲染
    if (Math.abs(newPosition - lastPosition) >= 1) {
      captchaSliderPosition.value = newPosition
      lastPosition = newPosition
    }

    animationFrameId = 0
  })
}

// 全局鼠标释放事件处理
const handleGlobalMouseUp = () => {
  isDragging.value = false

  // 取消可能存在的动画帧
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId)
    animationFrameId = 0
  }

  // 清除缓存的容器信息
  containerInfo = null

  // 移除全局鼠标事件监听
  document.removeEventListener('mousemove', handleGlobalMouseMove)
  document.removeEventListener('mouseup', handleGlobalMouseUp)

  // 使用requestAnimationFrame确保最后一次位置更新完成后再验证
  requestAnimationFrame(() => {
    setTimeout(() => {
      handleVerifyCaptcha()
    }, 1000) // 延迟1秒发送校验请求
  })
}

// 飞书扫码登录相关状态
const feishuLoginState = ref({
  isLoading: false,
  hasQrCode: false,
  sdkLoaded: false,
  showAuthIframe: false, // 控制是否显示授权iframe
})

// 保存飞书QRLogin实例引用和goto地址
let qrLoginInstance: any = null
let qrLoginGoto: string = ''

// 为window添加QRLogin类型定义
declare global {
  interface Window {
    QRLogin: any
  }
}

// 动态加载飞书二维码 SDK 脚本
const loadFeishuQRSDK = async () => {
  if (feishuLoginState.value.sdkLoaded) {
    initQRLogin()
    return
  }

  return new Promise<void>((resolve) => {
    const script = document.createElement('script')
    script.src =
      'https://lf-package-cn.feishucdn.com/obj/feishu-static/lark/passport/qrcode/LarkSSOSDKWebQRCode-1.0.3.js'
    script.onload = () => {
      feishuLoginState.value.sdkLoaded = true
      initQRLogin()
      resolve()
    }
    document.body.appendChild(script)
  })
}

// 初始化二维码登录实例
const initQRLogin = () => {
  // 检查容器是否存在
  const container = document.getElementById('login_container')
  if (!container) {
    console.warn('飞书二维码容器不存在，无法初始化')
    return
  }

  const clientId = 'cli_a9b81d3fcaf81bc1' // 替换为开发者后台的 App ID
  // 修改重定向地址为我们新创建的授权页面
  const redirectUri = encodeURIComponent(`${window.location.origin}/user/lark/auth`) // 需与开发者后台配置一致
  const state = Math.random().toString(36).substr(2) // 随机字符串防 CSRF

  // 清空容器内容，避免重复创建二维码实例
  container.innerHTML = ''

  // 构建goto地址并保存
  qrLoginGoto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code&state=${state}`

  // 重置状态
  feishuLoginState.value.showAuthIframe = false

  // 创建 QRLogin 实例
  qrLoginInstance = window.QRLogin({
    id: 'login_container', // 二维码容器的 DOM ID
    goto: qrLoginGoto,
    width: '300',
    height: '300',
  })

  // 确保只添加一个事件监听器
  window.removeEventListener('message', handleFeishuMessage)
  window.addEventListener('message', handleFeishuMessage)

  // 添加iframe消息监听
  window.removeEventListener('message', handleIframeMessage)
  window.addEventListener('message', handleIframeMessage)

  feishuLoginState.value.hasQrCode = true
}

// 飞书消息事件处理函数
const handleFeishuMessage = (event) => {
  // 检查事件数据是否有效
  if (event.data && event.data.tmp_code && qrLoginGoto) {
    // 验证事件来源和数据（如果实例存在）
    if (qrLoginInstance && qrLoginInstance.matchOrigin && qrLoginInstance.matchData) {
      if (!qrLoginInstance.matchOrigin(event.origin) || !qrLoginInstance.matchData(event.data)) {
        return
      }
    }

    const tmpCode = event.data.tmp_code
    // 拼接 tmp_code 到 goto 地址，显示内嵌授权iframe
    qrLoginGoto = `${qrLoginGoto}&tmp_code=${tmpCode}`
    feishuLoginState.value.showAuthIframe = true
  }
}

// 处理来自iframe的消息
const handleIframeMessage = (event: MessageEvent) => {
  // 验证消息来源
  if (event.origin !== window.location.origin) {
    return
  }

  // 检查是否是登录成功消息
  if (event.data && event.data.type === 'lark-login-success') {
    // 登录成功，关闭登录弹窗
    message.success('飞书登录成功')
    setTimeout(() => {
      handleClose()
    }, 2000)
  } else if (event.data && event.data.type === 'lark-login-failed') {
    // 登录失败，显示错误信息
    message.error(event.data.error || '飞书登录失败')
    // 重置状态，显示二维码
    feishuLoginState.value.showAuthIframe = false
    // 重新生成二维码
    setTimeout(() => {
      handleFeishuQRLogin()
    }, 1000)
  }
}

// 获取飞书登录二维码
const handleFeishuQRLogin = async () => {
  feishuLoginState.value.isLoading = true
  try {
    await loadFeishuQRSDK()
  } catch (error) {
    message.error('获取二维码失败')
  } finally {
    feishuLoginState.value.isLoading = false
  }
}

// 绘制波浪
let animationId: number
const drawWaves = (ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement) => {
  let time = 0

  const draw = () => {
    // 获取实际的canvas尺寸
    const width = canvas.width
    const height = canvas.height

    // 清空画布
    ctx.clearRect(0, 0, width, height)

    // 设置渐变
    const gradient = ctx.createLinearGradient(0, 0, 0, height)
    gradient.addColorStop(0, 'rgba(24, 144, 255, 0.8)')
    gradient.addColorStop(1, 'rgba(24, 144, 255, 0.1)')

    // 绘制第一条波浪
    ctx.beginPath()
    ctx.moveTo(0, height * 0.5)
    for (let x = 0; x <= width; x += 5) {
      const y =
        height * 0.5 +
        Math.sin((x / width) * Math.PI * 2 + time * 0.05) * 15 +
        Math.sin((x / width) * Math.PI + time * 0.02) * 10
      ctx.lineTo(x, y)
    }
    ctx.lineTo(width, height)
    ctx.lineTo(0, height)
    ctx.fillStyle = gradient
    ctx.fill()

    // 绘制第二条波浪
    ctx.beginPath()
    ctx.moveTo(0, height * 0.55)
    for (let x = 0; x <= width; x += 5) {
      const y =
        height * 0.55 +
        Math.sin((x / width) * Math.PI * 2 + time * 0.04) * 18 +
        Math.sin((x / width) * Math.PI + time * 0.03) * 12
      ctx.lineTo(x, y)
    }
    ctx.lineTo(width, height)
    ctx.lineTo(0, height)
    ctx.fillStyle = 'rgba(24, 144, 255, 0.3)'
    ctx.fill()

    time++
    animationId = requestAnimationFrame(draw)
  }

  draw()
}

// 初始化Canvas动画
const initWaveAnimation = async () => {
  await nextTick()
  const canvas = document.querySelector('.login-wave-canvas') as HTMLCanvasElement
  if (canvas) {
    // 设置canvas的实际像素大小
    const rect = canvas.getBoundingClientRect()
    canvas.width = rect.width
    canvas.height = rect.height

    const ctx = canvas.getContext('2d')
    if (ctx) {
      drawWaves(ctx, canvas)
    }
  }
}

onMounted(() => {
  initWaveAnimation()
  // 移除这里的loadFeishuQRSDK调用，只在需要显示二维码时才加载
})

watch(
  () => props.visible,
  (newVal) => {
    if (newVal) {
      initWaveAnimation()
      // 弹窗打开时，如果当前是飞书登录标签，自动加载二维码
      if (activeTab.value === 'feishu') {
        // 延迟加载，确保DOM已渲染
        setTimeout(() => {
          handleFeishuQRLogin()
        }, 100)
      }
    }
  },
)

// 监听当前激活的标签页，如果切换到飞书登录且弹窗可见，则自动加载二维码
watch(activeTab, (newTab) => {
  if (newTab === 'feishu' && props.visible) {
    // 延迟加载，确保DOM已渲染
    setTimeout(() => {
      handleFeishuQRLogin()
    }, 100)
  } else if (newTab !== 'feishu') {
    // 切换到其他标签页时，清理二维码资源
    window.removeEventListener('message', handleFeishuMessage)
    window.removeEventListener('message', handleIframeMessage)
    const container = document.getElementById('login_container')
    if (container) {
      container.innerHTML = ''
    }
    qrLoginInstance = null
    qrLoginGoto = ''
    feishuLoginState.value.showAuthIframe = false
  }
})

// 组件卸载时清理资源
onUnmounted(() => {
  // 移除消息事件监听
  window.removeEventListener('message', handleFeishuMessage)
  window.removeEventListener('message', handleIframeMessage)

  // 清空二维码容器
  const container = document.getElementById('login_container')
  if (container) {
    container.innerHTML = ''
  }

  // 清理实例引用和goto地址
  qrLoginInstance = null
  qrLoginGoto = ''

  // 清理验证码计时器
  if (timer) {
    clearInterval(timer)
    timer = null
  }
})
</script>

<template>
  <Modal
    :open="props.visible"
    :footer="null"
    @cancel="handleClose"
    width="420px"
    centered
    :mask-style="{ backdropFilter: 'blur(5px)' }"
  >
    <!-- 波浪背景 canvas -->
    <canvas class="login-wave-canvas"></canvas>

    <div class="login-modal-content">
      <h1 class="login-title">GoshowMall</h1>
      <p class="login-subtitle">欢迎登录</p>

      <Tabs v-model:activeKey="activeTab" animated class="login-tabs" size="small">
        <Tabs.TabPane key="account" tab="手机号登录">
          <Form layout="vertical">
            <Form.Item label="手机号" required>
              <Input
                v-model:value="accountForm.mobile"
                placeholder="请输入手机号"
                size="large"
                allow-clear
              >
                <template #prefix>
                  <PhoneOutlined />
                </template>
              </Input>
            </Form.Item>

            <Form.Item label="密码" required>
              <Input.Password
                v-model:value="accountForm.password"
                placeholder="请输入密码"
                size="large"
              >
                <template #prefix>
                  <LockOutlined />
                </template>
              </Input.Password>
            </Form.Item>

            <Form.Item>
              <Button
                type="primary"
                size="large"
                block
                :loading="isLoading"
                @click="handleAccountLogin"
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Tabs.TabPane>

        <!-- 验证码登录 -->
        <Tabs.TabPane key="code" tab="验证码登录">
          <Form layout="vertical">
            <Form.Item label="手机号" required>
              <Input
                v-model:value="codeForm.phone"
                placeholder="请输入手机号"
                size="large"
                allow-clear
              >
                <template #prefix>
                  <PhoneOutlined />
                </template>
              </Input>
            </Form.Item>

            <Form.Item label="验证码" required>
              <Space style="width: 100%; gap: 8px">
                <Input
                  v-model:value="codeForm.code"
                  placeholder="请输入验证码"
                  size="large"
                  style="flex: 1"
                  allow-clear
                />
                <Button
                  type="default"
                  size="large"
                  :disabled="countDown > 0"
                  @click="handleGetCode"
                  style="width: 120px; text-align: center; min-width: 120px"
                >
                  {{ countDown > 0 ? `${countDown}s` : '获取验证码' }}
                </Button>
              </Space>
            </Form.Item>

            <Form.Item>
              <Button
                type="primary"
                size="large"
                block
                :loading="isLoading"
                @click="handleCodeLogin"
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Tabs.TabPane>

        <!-- 飞书扫码登录 -->
        <Tabs.TabPane key="feishu" tab="飞书扫码">
          <div class="feishu-login-container">
            <!-- 二维码容器 -->
            <div v-if="!feishuLoginState.showAuthIframe" style="text-align: center">
              <div
                id="login_container"
                style="
                  width: 250px;
                  height: 250px;
                  margin: 0 auto;
                  display: flex;
                  align-items: center;
                  justify-content: center;
                "
              ></div>

              <!-- 加载中状态 -->
              <div
                v-if="feishuLoginState.isLoading"
                style="text-align: center; margin-top: 16px; color: #1890ff"
              >
                正在加载二维码...
              </div>
            </div>

            <!-- 授权iframe容器 -->
            <div v-else class="feishu-auth-iframe-container">
              <iframe
                :src="qrLoginGoto"
                class="feishu-auth-iframe"
                frameborder="0"
                allow="autoplay; fullscreen; clipboard-read; clipboard-write"
              ></iframe>
            </div>

            <p style="text-align: center; color: #666; margin-top: 16px">使用飞书扫描二维码登录</p>
          </div>
        </Tabs.TabPane>
      </Tabs>

      <!-- 底部链接 -->
      <div class="login-footer">
        <span>没有账号？</span>
        <a href="#" @click.prevent>立即注册</a>
        <span class="divider">|</span>
        <a href="#" @click.prevent>忘记密码？</a>
      </div>
    </div>
  </Modal>

  <!-- 滑块验证码弹窗 -->
  <Modal
    v-model:open="captchaModal"
    title="人机正常验证"
    :footer="null"
    @close="handleCloseCaptchaModal"
    centered
    width="350px"
  >
    <div class="captcha-container">
      <!-- 滑块验证码图片 -->
      <div class="captcha-image-wrapper" v-if="captchaData">
        <!-- 背景图 -->
        <img :src="captchaData.ImageBase64" alt="验证码" class="captcha-image" />
        <!-- 拼图 -->
        <img
          v-if="captchaData.TitleImageBase64"
          :src="captchaData.TitleImageBase64"
          alt="拼图"
          class="captcha-puzzle"
          :style="{
            left: captchaSliderPosition + 'px',
            top: captchaData.TitleY + 'px',
          }"
        />
      </div>
      <!-- 滑块轨道 -->
      <div class="captcha-slider-container" ref="sliderContainer">
        <div
          class="captcha-slider-handle"
          :class="{ dragging: isDragging }"
          :style="{ left: captchaSliderPosition + 'px' }"
          @mousedown="handleSliderMouseDown"
        >
          <span class="slider-icon">→</span>
        </div>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
:deep(.ant-modal-content) {
  position: relative;
  overflow: hidden;
  background: transparent;
}

:deep(.ant-modal-body) {
  position: relative;
  padding: 0;
  overflow: hidden;
  background: #fff;
}

.login-wave-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 120px;
  z-index: 0;
}

.login-modal-content {
  padding: 20px 24px;
  position: relative;
  z-index: 1;
}

.login-title {
  text-align: center;
  font-size: 28px;
  font-weight: bold;
  color: #1890ff;
  margin: 0 0 8px 0;
}

.login-subtitle {
  text-align: center;
  color: #666;
  margin: 0 0 24px 0;
  font-size: 14px;
}

.login-tabs {
  margin-bottom: 16px;
}

/* 简化Tab样式 */
:deep(.ant-tabs-nav) {
  margin: 16px !important;
  padding: 0 !important;
}

:deep(.ant-tabs-nav-wrap) {
  overflow: visible !important;
}

:deep(.ant-tabs-nav::before) {
  display: none !important;
}

:deep(.ant-tabs-nav-list) {
  width: 100% !important;
  justify-content: space-around;
}

:deep(.ant-tabs-tab) {
  margin: 4px !important;
  font-size: 14px !important;
  font-weight: 500 !important;
  flex: 0 0 auto !important;
}

:deep(.ant-tabs-tab-active .ant-tabs-tab-btn) {
  margin: 4px !important;
  color: #1890ff !important;
  font-weight: 600 !important;
}

:deep(.ant-tabs-ink-bar) {
  height: 3px !important;
  background-color: #1890ff !important;
  border-radius: 2px;
  bottom: -3px !important; /* 使下划线离文字编距壳异的三像姓等的上面空隙 */
  transition: all 0.3s ease !important;
}

/* 控制Tab内容滑动 */
:deep(.ant-tabs-content-holder) {
  /* 使用CSS transitions并限制动画在弹窗范围内 */
  overflow: hidden;
  position: relative;
}

:deep(.ant-tabs-tabpane) {
  animation: none !important;
  /* 禁用Ant Design默认动画 */
}

:deep(.ant-tabs-tab-active) {
  animation: slideIn 0.3s ease-in-out forwards;
}

@keyframes slideIn {
  from {
    opacity: 0.8;
    transform: translateX(10px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.feishu-login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  min-height: 300px;
}

/* 飞书二维码容器样式 */
#login_container {
  border: none !important;
  width: 250px !important;
  height: 250px !important;
  margin: 0 auto !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* 使用深度选择器移除飞书SDK生成的二维码边框 */
:deep(.qrcode-border) {
  border: none !important;
}

:deep(iframe) {
  border: none !important;
  margin: 0 !important;
  padding: 0 !important;
}

:deep(img) {
  border: none !important;
  margin: 0 auto !important;
}

/* 确保飞书SDK生成的内容完全居中 */
:deep(.qrcode-content) {
  margin: 0 auto !important;
  display: block !important;
}

.qrcode-placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 200px;
  height: 200px;
  border: 2px dashed #d9d9d9;
  border-radius: 8px;
  background-color: #fafafa;
}

.login-footer {
  text-align: center;
  margin-top: 24px;
  font-size: 14px;
  color: #666;
}

.login-footer a {
  color: #1890ff;
  text-decoration: none;
  margin: 0 4px;
  transition: color 0.3s;
}

.login-footer a:hover {
  color: #40a9ff;
}

.divider {
  margin: 0 8px;
  color: #ddd;
}

/* 滑块验证码样式 */
.captcha-container {
  display: flex;
  flex-direction: column;
  width: 100%;
  gap: 16px;
}

.captcha-image-wrapper {
  position: relative;
  overflow: hidden;
  border-radius: 4px;
  /* 添加GPU加速 */
  transform: translateZ(0);
  will-change: transform;
}

.captcha-image {
  position: relative;
  width: 100%;
}

.captcha-puzzle {
  position: absolute;
  z-index: 2;
  transition: none; /* 拖动时禁用过渡效果，提升性能 */
  pointer-events: none;
  width: auto;
  height: auto;
  will-change: left; /* 告诉浏览器我们将改变left属性，优化渲染 */
  /* 添加GPU加速 */
  transform: translateZ(0);
}

.captcha-slider-container {
  position: relative;
  width: 100%;
  height: 50px;
  background-color: #f7f9fa;
  border-radius: 4px;
  margin-bottom: 16px;
  border: 1px solid #e0e0e0;
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.05);
  /* 添加GPU加速 */
  transform: translateZ(0);
}

.captcha-slider-handle {
  position: absolute;
  top: 50%;
  left: 0;
  width: 50px;
  height: 50px;
  background-color: #fff;
  border-radius: 4px;
  cursor: grab;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  font-weight: bold;
  transition: all 0.05s ease-out;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
  user-select: none;
  z-index: 10;
  transform: translateY(-50%) translateZ(0); /* 合并transform属性并添加GPU加速 */
  border: 1px solid #e0e0e0;
  will-change: left, background-color, box-shadow; /* 优化渲染 */
}

.captcha-slider-handle:hover {
  background-color: #f0f8ff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.captcha-slider-handle.dragging {
  cursor: grabbing;
  background-color: #e6f7ff;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.25);
}

.slider-icon {
  font-size: 25px;
  color: #1890ff;
}

/* 飞书登录相关样式 */
.feishu-login-container {
  position: relative;
  width: 100%;
  min-height: 300px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.feishu-auth-iframe-container {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
}

.feishu-auth-iframe {
  width: 100%;
  height: 100%;
  border: none;
}
</style>
