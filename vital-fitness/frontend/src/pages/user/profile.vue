<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">个人资料</text>
			<view class="nav-placeholder"></view>
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
	import { useUserStore } from '../../store'

	export default {
		data() {
			return {
				saving: false,
				form: { nickname: '', gender: 0, height: '', weight: '' }
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
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
					const userStore = useUserStore()
					const updated = { ...userStore.userInfo, nickname: this.form.nickname, gender: this.form.gender }
					userStore.userInfo = updated
					uni.setStorageSync('userInfo', JSON.stringify(updated))
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	@include card;
	margin-bottom: $spacing-md;
	.card-label { @include card-label; }
}

.input-row {
	@include input-field;
	.field-input { flex: 1; font-size: $font-body; color: $color-label; }
	.input-suffix { font-size: $font-subhead; color: $color-label-quaternary; }
}

.segment { @include segment-control; }

.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}
</style>
