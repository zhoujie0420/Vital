<template>
	<view class="container">
		<u-navbar title="个人资料" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="form-container">
			<view class="form-item">
				<text class="label">昵称</text>
				<input v-model="form.nickname" placeholder="输入昵称" class="input-field" />
			</view>
			<view class="form-item">
				<text class="label">性别</text>
				<view class="gender-tabs">
					<text class="gender-tab" :class="{ active: form.gender === 1 }" @tap="form.gender = 1">男</text>
					<text class="gender-tab" :class="{ active: form.gender === 2 }" @tap="form.gender = 2">女</text>
					<text class="gender-tab" :class="{ active: form.gender === 0 }" @tap="form.gender = 0">保密</text>
				</view>
			</view>
			<view class="form-item">
				<text class="label">身高 (cm)</text>
				<input v-model="form.height" type="digit" placeholder="输入身高" class="input-field" />
			</view>
			<view class="form-item">
				<text class="label">体重 (kg)</text>
				<input v-model="form.weight" type="digit" placeholder="输入体重" class="input-field" />
			</view>
			<button class="save-btn" type="primary" :loading="saving" @tap="save">
				{{ saving ? '保存中...' : '保存' }}
			</button>
		</view>
	</view>
</template>

<script>
	import { getProfile } from '../../api/user'
	import { put } from '../../utils/request'

	export default {
		data() {
			return {
				saving: false,
				form: { nickname: '', gender: 0, height: '', weight: '' }
			}
		},
		onLoad() { this.loadProfile() },
		methods: {
			goBack() { uni.navigateBack() },
			async loadProfile() {
				try {
					const res = await getProfile()
					const d = res.data || {}
					this.form.nickname = d.nickname || ''
					this.form.gender = d.gender || 0
					this.form.height = d.height || ''
					this.form.weight = d.weight || ''
				} catch (e) {}
			},
			async save() {
				this.saving = true
				try {
					await put('/users/profile', {
						nickname: this.form.nickname,
						gender: this.form.gender,
						height: parseFloat(this.form.height) || 0,
						weight: parseFloat(this.form.weight) || 0
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.form-container {
	.form-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 15rpx;
		.label { font-size: 28rpx; color: #666; margin-bottom: 15rpx; display: block; }
		.input-field { border: 1rpx solid #eee; border-radius: 10rpx; padding: 20rpx; font-size: 28rpx; }
		.gender-tabs { display: flex; gap: 15rpx;
			.gender-tab {
				flex: 1; text-align: center; padding: 15rpx; border-radius: 10rpx; background: #f5f5f5; font-size: 28rpx; color: #666;
				&.active { background: #3c9cff; color: white; }
			}
		}
	}
	.save-btn { margin-top: 20rpx; border-radius: 15rpx; font-size: 32rpx; }
}
</style>
