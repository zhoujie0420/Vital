import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: ['src/**/__tests__/**/*.test.js'],
    environment: 'node',
    setupFiles: ['src/__tests__/setup.js'],
    poolOptions: {
      threads: { singleThread: true }
    }
  }
})
