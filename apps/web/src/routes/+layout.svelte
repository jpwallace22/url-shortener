<script>
  import AnimatedBackground from '../lib/common/AnimatedBackground.svelte';
  import '../app.css';
</script>

<body class="dark h-svh grid place-items-center w-full">
  <AnimatedBackground />
  <div class="w-full grid gap-16 max-w-lg px-3">
    <slot />
  </div>
</body>
