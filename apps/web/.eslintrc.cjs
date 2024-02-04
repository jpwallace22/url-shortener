module.exports = {
	extends: ['@repo/eslint-config/index.js'],
	rules: {
		'@typescript-eslint/no-unused-vars': [
			'warn',
			{
				argsIgnorePattern: '^_',
				varsIgnorePattern: '^$$(Props|Events|Slots|Generic)$'
			}
		]
	}
};
