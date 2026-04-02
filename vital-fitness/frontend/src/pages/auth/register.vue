<template>
	<view class="container">
		<u-navbar title="注册账号" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="register-form">
			<u-form :model="registerForm" ref="registerFormRef">
				<u-form-item>
					<u-input
						v-model="registerForm.username"
						placeholder="用户名（3-20个字符）"
						prefix-icon="account"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<u-form-item>
					<u-input
						v-model="registerForm.email"
						placeholder="邮箱地址"
						type="text"
						prefix-icon="email"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<u-form-item>
					<u-input
						v-model="registerForm.password"
						placeholder="密码（至少6位）"
						type="password"
						prefix-icon="lock"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<u-form-item>
					<u-input
						v-model="registerForm.confirmPassword"
						placeholder="确认密码"
						type="password"
						prefix-icon="lock"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<u-button
					type="primary"
					@click="handleRegister"
					:loading="loading"
					:disabled="loading"
					style="margin-top: 40rpx;"
				>
					{{ loading ? '注册中...' : '注册' }}
				</u-button>
			</u-form>

			<view class="login-link">
				<text class="login-text">已有账号？</text>
				<text class="login-btn" @click="toLogin">立即登录</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { register } from '../../api/user'

	export default {
		data() {
			return {
				loading: false,
				registerForm: {
					username: '',
					email: '',
					password: '',
					confirmPassword: ''
				}
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			validate() {
				const { username, email, password, confirmPassword } = this.registerForm

				if (!username || username.length < 3 || username.length > 20) {
					uni.showToast({ title: '用户名长度为3-20个字符', icon: 'none' })
					return false
				}

				const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
				if (!email || !emailRegex.test(email)) {
					uni.showToast({ title: '请输入正确的邮箱地址', icon: 'none' })
					return false
				}

				if (!password || password.length < 6) {
					uni.showToast({ title: '密码长度不能少于6位', icon: 'none' })
					return false
				}

				if (password !== confirmPassword) {
					uni.showToast({ title: '两次输入的密码不一致', icon: 'none' })
					return false
				}

				return true
			},

			async handleRegister() {
				if (!this.validate()) return

				this.loading = true
				try {
					await register({
						username: this.registerForm.username,
						email: this.registerForm.email,
						password: this.registerForm.password
					})

					uni.showToast({ title: '注册成功', icon: 'success' })
					setTimeout(() => {
						uni.redirectTo({ url: '/pages/auth/login' })
					}, 1500)
				} catch (err) {
					// request.js 已处理 toast
				} finally {
					this.loading = false
				}
			},

			toLogin() {
				uni.navigateBack()
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
	min-height: 100vh;
	background: #f5f5f5;
}

.register-form {
	background: white;
	border-radius: 20rpx;
	padding: 60rpx 40rpx;
	margin: 30rpx;

	.login-link {
		text-align: center;
		margin-top: 40rpx;

		.login-text { font-size: 26rpx; color: #999; margin-right: 10rpx; }
		.login-btn { font-size: 26rpx; color: #3c9cff; font-weight: bold; }
	}
}
</style>
