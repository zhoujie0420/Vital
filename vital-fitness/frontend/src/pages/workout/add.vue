<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录训练</text>
			<text class="nav-placeholder"></text>
		</view>

		<view class="card">
			<text class="card-label">选择动作</text>
			<scroll-view scroll-x class="cat-scroll">
				<text v-for="cat in categories" :key="cat" class="cat-pill"
					:class="{ active: currentCategory === cat }" @tap="currentCategory = cat">
					{{ cat }}
				</text>
			</scroll-view>

			<view class="exercise-grid">
				<view v-for="ex in filteredExercises" :key="ex.id" class="exercise-chip"
					:class="{ selected: selectedExercise && selectedExercise.id === ex.id }"
					@tap="selectExercise(ex)">
					<text class="exercise-name">{{ ex.name }}</text>
				</view>
			</view>
		</view>

		<view class="card" v-if="selectedExercise">
			<text class="card-label">重量</text>
			<view class="weight-display">
				<text class="weight-btn" @tap="adjustWeight(-2.5)">−</text>
				<view class="weight-center">
					<input type="digit" v-model="weight" class="weight-num" placeholder="0" />
					<text class="weight-unit">kg</text>
				</view>
				<text class="weight-btn" @tap="adjustWeight(2.5)">+</text>
			</view>
			<view class="quick-row">
				<text class="quick-pill" v-for="w in [20, 40, 60, 80, 100]" :key="w"
					:class="{ active: parseFloat(weight) === w }" @tap="weight = w">{{ w }}kg</text>
			</view>
		</view>

		<view class="bottom-action" v-if="selectedExercise">
			<view class="save-btn" :class="{ loading: saving }" @tap="saveWorkout">
				<text>{{ saving ? '保存中...' : '保存记录' }}</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed, onMounted } from 'vue'
	import { getExercises, createWorkout } from '../../api/workout'

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const saving = ref(false)
	const categories = ['全部', '胸部', '背部', '腿部', '肩部', '手臂', '核心']
	const currentCategory = ref('全部')
	const exercises = ref([])
	const selectedExercise = ref(null)
	const weight = ref('')

	const filteredExercises = computed(() => {
		if (currentCategory.value === '全部') return exercises.value
		return exercises.value.filter(ex => ex.category === currentCategory.value)
	})

	const goBack = () => uni.navigateBack()
	const selectExercise = (ex) => { selectedExercise.value = ex }
	const adjustWeight = (d) => {
		const c = parseFloat(weight.value) || 0
		if (c + d >= 0) weight.value = c + d
	}

	const saveWorkout = async () => {
		if (!weight.value || parseFloat(weight.value) <= 0) {
			uni.showToast({ title: '请输入重量', icon: 'none' }); return
		}
		saving.value = true
		try {
			await createWorkout({
				exercise_id: selectedExercise.value.id,
				workout_date: new Date().toISOString(),
				weight: parseFloat(weight.value), sets: 1, reps: 1
			})
			uni.showToast({ title: '保存成功', icon: 'success' })
			setTimeout(() => uni.navigateBack(), 1000)
		} catch (e) {} finally { saving.value = false }
	}

	onMounted(async () => {
		try {
			const res = await getExercises('')
			exercises.value = res.data || []
		} catch (e) {}
	})
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.card-label {
		display: block;
		font-size: 26rpx;
		font-weight: 600;
		color: #8e8e93;
		text-transform: uppercase;
		letter-spacing: 1rpx;
		margin-bottom: 20rpx;
	}
}

.cat-scroll {
	white-space: nowrap;
	margin-bottom: 20rpx;

	.cat-pill {
		display: inline-block;
		padding: 10rpx 24rpx;
		border-radius: 20rpx;
		background: #f2f2f7;
		font-size: 26rpx;
		color: #636366;
		margin-right: 12rpx;
		font-weight: 500;
		transition: all 0.2s ease;

		&.active {
			background: #007aff;
			color: #fff;
			box-shadow: 0 4rpx 12rpx rgba(0, 122, 255, 0.3);
		}
		&:active { transform: scale(0.95); }
	}
}

.exercise-grid {
	display: flex;
	flex-wrap: wrap;
	gap: 12rpx;

	.exercise-chip {
		padding: 16rpx 28rpx;
		border-radius: 16rpx;
		background: #f2f2f7;
		border: 2rpx solid transparent;
		transition: all 0.2s ease;

		&.selected {
			background: rgba(0, 122, 255, 0.08);
			border-color: #007aff;
			box-shadow: 0 2rpx 12rpx rgba(0, 122, 255, 0.15);
		}
		&:active { transform: scale(0.95); }

		.exercise-name {
			font-size: 28rpx;
			color: #1c1c1e;
			font-weight: 500;
		}
	}
}

.weight-display {
	display: flex;
	align-items: center;
	justify-content: center;
	padding: 20rpx 0;

	.weight-btn {
		width: 72rpx;
		height: 72rpx;
		line-height: 68rpx;
		text-align: center;
		background: #f2f2f7;
		border-radius: 50%;
		font-size: 40rpx;
		color: #007aff;
		font-weight: 300;

		&:active { background: #e5e5ea; transform: scale(0.92); }
	}

	.weight-center {
		display: flex;
		align-items: baseline;
		margin: 0 32rpx;

		.weight-num {
			width: 160rpx;
			text-align: center;
			font-size: 60rpx;
			font-weight: 700;
			color: #1c1c1e;
			letter-spacing: -2rpx;
		}
		.weight-unit {
			font-size: 30rpx;
			color: #8e8e93;
			margin-left: 4rpx;
		}
	}
}

.quick-row {
	display: flex;
	justify-content: center;
	flex-wrap: wrap;
	gap: 12rpx;
	margin-top: 16rpx;

	.quick-pill {
		padding: 12rpx 28rpx;
		background: #f2f2f7;
		border-radius: 20rpx;
		font-size: 26rpx;
		color: #636366;
		font-weight: 500;
		transition: all 0.2s ease;

		&.active { background: #007aff; color: #fff; }
		&:active { transform: scale(0.95); opacity: 0.7; }
	}
}

.bottom-action {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	padding: 20rpx 32rpx;
	padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
	background: rgba(242, 242, 247, 0.9);
	backdrop-filter: blur(20px);
	-webkit-backdrop-filter: blur(20px);

	.save-btn {
		background: #007aff;
		color: #fff;
		text-align: center;
		padding: 28rpx 0;
		border-radius: 16rpx;
		font-size: 32rpx;
		font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);
		transition: all 0.15s ease;

		&:active { transform: scale(0.98); opacity: 0.9; }
		&.loading { opacity: 0.6; }
	}
}
</style>
