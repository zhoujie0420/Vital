<template>
	<view class="page">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">个人资料</text>
			<text class="nav-placeholder"></text>
		</view>

		<view class="card">
			<text class="card-label">昵称</text>
			<view class="input-row">
				<input v-model="form.nickname" placeholder="输入昵称" class="field-input" />
			</view>
		</view>

		<view class="card">
			<text class="card-label">性别</text>
			<view class="segment">
				<text class="segment-item" :class="{ active: form.gender === 1 }" @tap="form.gender = 1">男</text>
				<text class="segment-item" :class="{ active: form.gender === 2 }" @tap="form.gender = 2">女</text>
				<text class="segment-item" :class="{ active: form.gender === 0 }" @tap="form.gender = 0">保密</text>
			</view>
		</view>

		<view class="card">
			<text class="card-label">身高</text>
			<view class="input-row">
				<input v-model="form.height" type="digit" placeholder="输入身高" class="field-input" />
				<text class="input-suffix">cm</text>
			</view>
		</view>

		<view class="card">
			<text class="card-label">体重</text>
			<view class="input-row">
				<input v-model="form.weight" type="digit" placeholder="输入体重" class="field-input" />
				<text class="input-suffix">kg</text>
			</view>
		</view>

		<view class="bottom-action">
			<view class="save-btn" :class="{ loading: saving }" @tap="save">
				<text>{{ saving ? '保存中...' : '保存' }}</text>
			</view>
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
.page {
	padding: 0 32rpx; padding-top: 120rpx; padding-bottom: 160rpx;
	min-height: 100vh; background: #f2f2f7;
}

.nav-bar {
	display: flex; align-items: center; justify-content: space-between; margin-bottom: 28rpx;
	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	background: #fff; border-radius: 20rpx; padding: 28rpx 32rpx; margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	.card-label {
		display: block; font-size: 26rpx; font-weight: 600; color: #8e8e93;
		text-transform: uppercase; letter-spacing: 1rpx; margin-bottom: 16rpx;
	}
}

.input-row {
	display: flex; align-items: center; background: #f2f2f7; border-radius: 16rpx; padding: 20rpx;
	.field-input { flex: 1; font-size: 30rpx; color: #1c1c1e; }
	.input-suffix { font-size: 28rpx; color: #8e8e93; }
}

.segment {
	display: flex; background: #f2f2f7; border-radius: 16rpx; padding: 4rpx;
	.segment-item {
		flex: 1; text-align: center; padding: 16rpx 0; font-size: 28rpx;
		font-weight: 500; color: #636366; border-radius: 14rpx; transition: all 0.2s;
		&.active {
			background: #fff; color: #1c1c1e; font-weight: 600;
			box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.08);
		}
	}
}

.bottom-action {
	position: fixed; bottom: 0; left: 0; right: 0;
	padding: 20rpx 32rpx; padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
	background: rgba(242, 242, 247, 0.9); backdrop-filter: blur(20px);
	.save-btn {
		background: #007aff; color: #fff; text-align: center; padding: 28rpx 0;
		border-radius: 16rpx; font-size: 32rpx; font-weight: 600;
		&:active { opacity: 0.8; }
		&.loading { opacity: 0.6; }
	}
}
</style>
