<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录训练</text>
			<view class="nav-placeholder"></view>
		</view>

		<view class="card">
			<text class="card-label">选择动作</text>
			<scroll-view scroll-x class="cat-scroll" :show-scrollbar="false">
				<view class="cat-bar">
					<text v-for="cat in categories" :key="cat" class="cat-pill"
						:class="{ active: currentCategory === cat }" @tap="currentCategory = cat">
						{{ cat }}
					</text>
				</view>
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
				<view class="weight-btn" @tap="adjustWeight(-2.5)">
					<text class="weight-btn-text">−</text>
				</view>
				<view class="weight-center">
					<input type="digit" v-model="weight" class="weight-num" placeholder="0" />
					<text class="weight-unit">kg</text>
				</view>
				<view class="weight-btn" @tap="adjustWeight(2.5)">
					<text class="weight-btn-text">+</text>
				</view>
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

// --- Category Scroll ---
.cat-scroll {
	white-space: nowrap;
	margin: 0 -#{$spacing-xl};
	padding: 0 $spacing-xl;
	margin-bottom: $spacing-md;

	.cat-bar {
		display: inline-flex;
		gap: $spacing-sm;
	}

	.cat-pill {
		display: inline-block;
		padding: $spacing-xs $spacing-lg;
		border-radius: $radius-full;
		background: $color-fill;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active {
			background: $color-label;
			color: #fff;
		}
		&:active { transform: scale(0.94); }
	}
}

// --- Exercise Grid ---
.exercise-grid {
	display: flex;
	flex-wrap: wrap;
	gap: $spacing-sm;

	.exercise-chip {
		padding: $spacing-md $spacing-lg;
		border-radius: $radius-md;
		background: $color-fill;
		border: 2rpx solid transparent;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.selected {
			background: $color-primary-light;
			border-color: $color-primary;
		}
		&:active { transform: scale(0.94); }

		.exercise-name {
			font-size: $font-subhead;
			color: $color-label;
			font-weight: 500;
		}
	}
}

// --- Weight Display ---
.weight-display {
	display: flex;
	align-items: center;
	justify-content: center;
	padding: $spacing-lg 0;

	.weight-btn {
		width: 72rpx;
		height: 72rpx;
		border-radius: 50%;
		background: $color-fill;
		display: flex;
		align-items: center;
		justify-content: center;
		@include press-effect;

		.weight-btn-text {
			font-size: 44rpx;
			color: $color-primary;
			font-weight: 300;
			line-height: 1;
		}
	}

	.weight-center {
		display: flex;
		align-items: baseline;
		margin: 0 $spacing-xl;

		.weight-num {
			width: 180rpx;
			text-align: center;
			font-size: $font-largeTitle;
			font-weight: 700;
			color: $color-label;
			letter-spacing: -2rpx;
			font-variant-numeric: tabular-nums;
		}
		.weight-unit {
			font-size: $font-subhead;
			color: $color-label-quaternary;
			margin-left: 4rpx;
		}
	}
}

.quick-row {
	display: flex;
	justify-content: center;
	flex-wrap: wrap;
	gap: $spacing-sm;
	margin-top: $spacing-md;

	.quick-pill {
		padding: $spacing-sm $spacing-lg;
		background: $color-fill;
		border-radius: $radius-full;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active { background: $color-label; color: #fff; }
		&:active { transform: scale(0.94); }
	}
}

// --- Bottom Action ---
.bottom-action {
	@include bottom-action-bar;

	.save-btn { @include primary-button; }
}
</style>
