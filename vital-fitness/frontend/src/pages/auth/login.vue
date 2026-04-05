<template>
	<view class="container">
		<view class="header-decoration">
			<view class="decoration-circle circle-1"></view>
			<view class="decoration-circle circle-2"></view>
		</view>

		<view class="logo-section">
			<text class="app-title">Vital Fitness</text>
			<text class="app-subtitle">记录每一个进步的瞬间</text>
		</view>

		<view class="login-section">
			<button class="wx-login-btn" :loading="loading" @tap="handleWxLogin">
				{{ loading ? '登录中...' : '微信一键登录' }}
			</button>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { useUserStore } from '../../store'
	import { wxLogin } from '../../api/user'

	const loading = ref(false)
	const userStore = useUserStore()

	const handleWxLogin = () => {
		loading.value = true
		uni.login({
			provider: 'weixin',
			success: async (loginRes) => {
				try {
					const res = await wxLogin(loginRes.code)
					userStore.saveLoginInfo({ token: res.data.token, user: res.data.user })
					uni.showToast({ title: '登录成功', icon: 'success' })
					setTimeout(() => uni.switchTab({ url: '/pages/index/index' }), 1000)
				} catch (err) {
					uni.showToast({ title: '登录失败', icon: 'none' })
				} finally { loading.value = false }
			},
			fail: () => {
				loading.value = false
				uni.showToast({ title: '微信登录取消', icon: 'none' })
			}
		})
	}
</script>

<style lang="scss" scoped>
.container {
	min-height: 100vh;
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	display: flex; flex-direction: column; justify-content: center; padding: 40rpx;
}
.header-decoration {
	.decoration-circle { position: absolute; border-radius: 50%; background: rgba(255,255,255,0.1); }
	.circle-1 { width: 200rpx; height: 200rpx; top: -100rpx; right: -100rpx; }
	.circle-2 { width: 300rpx; height: 300rpx; top: 200rpx; left: -150rpx; }
}
.logo-section {
	text-align: center; padding: 100rpx 0 80rpx;
	.app-title { font-size: 48rpx; font-weight: bold; color: white; display: block; margin-bottom: 20rpx; }
	.app-subtitle { font-size: 28rpx; color: rgba(255,255,255,0.8); }
}
.login-section { padding: 60rpx 40rpx;
	.wx-login-btn { background: #07c160; color: white; border: none; border-radius: 20rpx; font-size: 32rpx; padding: 24rpx 0; }
}
</style>
