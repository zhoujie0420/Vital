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

		<!-- 组数记录 — Hevy/Strong 风格 -->
		<view class="card" v-if="selectedExercise">
			<view class="exercise-header">
				<text class="exercise-title">{{ selectedExercise.name }}</text>
				<text class="exercise-cat">{{ selectedExercise.category }}</text>
			</view>

			<!-- 表头 -->
			<view class="set-table-header">
				<text class="col-set">组</text>
				<text class="col-prev">上次</text>
				<text class="col-weight">重量(kg)</text>
				<text class="col-reps">次数</text>
				<text class="col-action"></text>
			</view>

			<!-- 每组数据 -->
			<view v-for="(set, i) in sets" :key="i" class="set-row" :class="{ done: set.done }">
				<text class="col-set set-num">{{ i + 1 }}</text>
				<text class="col-prev set-prev">—</text>
				<view class="col-weight">
					<input type="digit" v-model="set.weight" class="set-input" placeholder="0"
						:class="{ filled: set.weight }" />
				</view>
				<view class="col-reps">
					<input type="number" v-model="set.reps" class="set-input" placeholder="0"
						:class="{ filled: set.reps }" />
				</view>
				<view class="col-action">
					<view class="check-btn" :class="{ checked: set.done }" @tap="set.done = !set.done">
						<text class="check-icon">✓</text>
					</view>
				</view>
			</view>

			<!-- 添加组 -->
			<view class="add-set-btn" @tap="addSet">
				<text class="add-set-text">+ 添加一组</text>
			</view>
		</view>

		<view class="bottom-action" v-if="selectedExercise">
			<view class="save-btn" :class="{ loading: saving }" @tap="saveWorkout">
				<text>{{ saving ? '保存中...' : '保存训练' }}</text>
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

	const sets = ref([
		{ weight: '', reps: '', done: false },
		{ weight: '', reps: '', done: false },
		{ weight: '', reps: '', done: false }
	])

	const addSet = () => {
		// 复制上一组的重量和次数
		const last = sets.value[sets.value.length - 1]
		sets.value.push({
			weight: last ? last.weight : '',
			reps: last ? last.reps : '',
			done: false
		})
	}

	const filteredExercises = computed(() => {
		if (currentCategory.value === '全部') return exercises.value
		return exercises.value.filter(ex => ex.category === currentCategory.value)
	})

	const goBack = () => uni.navigateBack()

	const selectExercise = (ex) => {
		selectedExercise.value = ex
		sets.value = [
			{ weight: '', reps: '', done: false },
			{ weight: '', reps: '', done: false },
			{ weight: '', reps: '', done: false }
		]
	}

	const saveWorkout = async () => {
		const validSets = sets.value.filter(s => parseFloat(s.weight) > 0 && parseInt(s.reps) > 0)
		if (validSets.length === 0) {
			uni.showToast({ title: '请至少填写一组数据', icon: 'none' }); return
		}
		saving.value = true
		try {
			// 保存每一组有效数据
			for (const s of validSets) {
				await createWorkout({
					exercise_id: selectedExercise.value.id,
					workout_date: new Date().toISOString(),
					weight: parseFloat(s.weight),
					sets: validSets.length,
					reps: parseInt(s.reps)
				})
			}
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

	.cat-bar { display: inline-flex; gap: $spacing-sm; }

	.cat-pill {
		display: inline-block;
		padding: $spacing-xs $spacing-lg;
		border-radius: $radius-full;
		background: $color-fill;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);

		&.active { background: $color-label; color: #fff; }
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
		transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);

		&.selected { background: $color-primary-light; border-color: $color-primary; }
		&:active { transform: scale(0.94); }

		.exercise-name { font-size: $font-subhead; color: $color-label; font-weight: 500; }
	}
}

// --- Exercise Header ---
.exercise-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: $spacing-lg;

	.exercise-title {
		font-size: $font-headline;
		font-weight: 700;
		color: $color-primary;
	}
	.exercise-cat {
		font-size: $font-caption1;
		color: $color-label-quaternary;
		font-weight: 500;
		background: $color-fill;
		padding: 4rpx $spacing-md;
		border-radius: $spacing-xs;
	}
}

// --- Set Table ---
.set-table-header {
	display: flex;
	align-items: center;
	padding: 0 0 $spacing-sm;
	border-bottom: 0.5rpx solid $color-separator;

	text {
		font-size: $font-caption2;
		font-weight: 600;
		color: $color-label-quaternary;
		text-transform: uppercase;
		letter-spacing: 1rpx;
	}
}

.col-set { width: 60rpx; text-align: center; }
.col-prev { width: 100rpx; text-align: center; }
.col-weight { flex: 1; text-align: center; }
.col-reps { flex: 1; text-align: center; }
.col-action { width: 64rpx; text-align: center; }

.set-row {
	display: flex;
	align-items: center;
	padding: $spacing-sm 0;
	border-bottom: 0.5rpx solid $color-separator;
	transition: background 0.2s ease;

	&.done {
		background: rgba(16, 185, 129, 0.04);
		border-radius: $radius-sm;
	}

	&:last-of-type {
		border-bottom: none;
	}

	.set-num {
		font-size: $font-subhead;
		font-weight: 700;
		color: $color-label-tertiary;
		text-align: center;
	}
	.set-prev {
		font-size: $font-caption1;
		color: $color-separator-opaque;
		text-align: center;
	}
}

.set-input {
	width: 100%;
	height: 64rpx;
	background: $color-fill;
	border-radius: $radius-sm;
	text-align: center;
	font-size: $font-subhead;
	font-weight: 600;
	color: $color-label;
	font-variant-numeric: tabular-nums;
	margin: 0 $spacing-xs;
	transition: all 0.2s ease;

	&.filled {
		background: $color-primary-light;
		color: $color-primary;
	}
}

.check-btn {
	width: 48rpx;
	height: 48rpx;
	border-radius: 50%;
	border: 2rpx solid $color-separator-opaque;
	display: flex;
	align-items: center;
	justify-content: center;
	margin: 0 auto;
	transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);

	.check-icon {
		font-size: $font-caption1;
		color: transparent;
		font-weight: 700;
	}

	&.checked {
		background: $color-primary;
		border-color: $color-primary;

		.check-icon { color: #fff; }
	}

	&:active { transform: scale(0.9); }
}

// --- Add Set ---
.add-set-btn {
	padding: $spacing-md 0 $spacing-xs;
	text-align: center;
	@include press-effect;

	.add-set-text {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
	}
}

// --- Bottom Action ---
.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}
</style>
