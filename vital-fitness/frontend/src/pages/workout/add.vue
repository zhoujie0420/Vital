<template>
	<view class="container">
		<view class="nav-bar">
			<text class="back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录训练</text>
		</view>

		<view class="form-item">
			<text class="label">选择动作</text>
			<view class="category-tabs">
				<text v-for="cat in categories" :key="cat" class="cat-tab"
					:class="{ active: currentCategory === cat }" @tap="currentCategory = cat">{{ cat }}</text>
			</view>
			<view class="exercise-grid">
				<view v-for="ex in filteredExercises" :key="ex.id" class="exercise-card"
					:class="{ selected: selectedExercise && selectedExercise.id === ex.id }"
					@tap="selectExercise(ex)">
					<text class="exercise-name">{{ ex.name }}</text>
				</view>
			</view>
		</view>

		<view class="form-item" v-if="selectedExercise">
			<text class="label">重量 (kg)</text>
			<view class="weight-input">
				<text class="weight-btn" @tap="adjustWeight(-2.5)">-</text>
				<input type="digit" v-model="weight" class="weight-value" placeholder="0" />
				<text class="weight-btn" @tap="adjustWeight(2.5)">+</text>
			</view>
			<view class="quick-weights">
				<text class="quick-btn" v-for="w in [20,40,60,80,100]" :key="w" @tap="weight = w">{{w}}kg</text>
			</view>
		</view>

		<button v-if="selectedExercise" class="save-btn" :loading="saving" @tap="saveWorkout">
			{{ saving ? '保存中...' : '保存记录' }}
		</button>
	</view>
</template>

<script setup>
	import { ref, computed, onMounted } from 'vue'
	import { getExercises, createWorkout } from '../../api/workout'

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
.container { padding: 20rpx; }
.nav-bar { display: flex; align-items: center; gap: 20rpx; margin-bottom: 20rpx;
	.back { font-size: 32rpx; color: #007aff; }
	.nav-title { font-size: 36rpx; font-weight: 700; color: #1c1c1e; }
}
.form-item {
	background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 20rpx;
	.label { font-size: 32rpx; font-weight: bold; color: #333; margin-bottom: 20rpx; display: block; }
	.category-tabs { display: flex; flex-wrap: wrap; gap: 15rpx; margin-bottom: 20rpx;
		.cat-tab { padding: 10rpx 24rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 24rpx; color: #666;
			&.active { background: #3c9cff; color: white; }
		}
	}
	.exercise-grid { display: flex; flex-wrap: wrap; gap: 15rpx;
		.exercise-card { padding: 20rpx 30rpx; border-radius: 15rpx; background: #f5f5f5; border: 2rpx solid transparent;
			&.selected { background: #e8f4ff; border-color: #3c9cff; }
			.exercise-name { font-size: 26rpx; color: #333; }
		}
	}
	.weight-input { display: flex; align-items: center; justify-content: center; margin: 30rpx 0;
		.weight-btn { width: 80rpx; height: 80rpx; line-height: 80rpx; text-align: center; background: #3c9cff; color: white; border-radius: 50%; font-size: 40rpx; }
		.weight-value { width: 200rpx; text-align: center; font-size: 60rpx; font-weight: bold; margin: 0 30rpx; }
	}
	.quick-weights { display: flex; justify-content: center; gap: 20rpx;
		.quick-btn { padding: 10rpx 24rpx; background: #f5f5f5; border-radius: 10rpx; font-size: 24rpx; color: #666; }
	}
}
.save-btn { border-radius: 15rpx; font-size: 32rpx; background: #007aff; color: white; }
</style>
