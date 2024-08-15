// uno.config.ts
import {
    defineConfig,
    presetAttributify,
    presetIcons,
    presetTypography,
    presetUno,
    presetWebFonts,
    transformerDirectives,
    transformerVariantGroup
} from 'unocss'

export default defineConfig({
    rules: [
        [/^m-(\d+)$/, ([, d]) => ({ margin: `${d as any / 4}rem` })],
        ['c-1', { color: 'red' }],
        ['bgc-w', { 'background-color': '#fff' }]
    ],
    shortcuts: [
        // ...
    ],
    theme: {
        colors: {
            // ...
            'veryCool': '#0000ff'
        }
    },
    presets: [
        presetUno(),
        presetAttributify(),
        presetIcons(),
        presetTypography(),
        presetWebFonts({
            provider: 'bunny',
            fonts: {
                'allura': 'Allura'
            }
        })
    ],
    transformers: [transformerDirectives(), transformerVariantGroup()]
})