<script setup lang="ts">
import FAQSection from '../components/FAQSection.vue'
</script>

<template>
  <div
    style="
      width: 100%;
      position: relative;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      color: white;
    "
  >
    <h1 style="margin-bottom: 60px">FAQ</h1>
    <div style="display: flex; flex-direction: column; width: 100%">
      <FAQSection
        title="What do these tests even measure?"
        description="These tests try to measure performance of pure backend logic meaning that the HTTP traffic itself is not included in the metric. While the algorithm section provide the best overview of how fast a language itself is, the other sections are also of interest since these represent real-life practical operations that get performed very frequently.
        The tests dont have the goal of showing off the best possible language implementation but the approach that would be usually used by devs in that context in a real-life setting. Since some libraries in a language might be implemented in a better and more performant way than 
        a library in another language, it sometimes is not a difference in language performance, but in library implementation performance (such as SQLite libraries of NodeJS and Go, where the NodeJS implementation is very performant)."
      />
      <FAQSection
        title="How are these tests conducted?"
        description="Each test is run 5 times and the average time taken is then measured. When possible an implementation with the standard library is used, but when not possible, the most popular and typically used 
      library for that use case is used. The backend logic is implemented in popular backend frameworks (namely Fiber for Go, Express.js for NodeJS and Django for Python)."
      />
      <FAQSection
        title="Some remarks regarding the Go implementation"
        description="The only external library used is one for SQLite operations."
      />
      <FAQSection
        title="Some remarks regarding the NodeJS implementation"
        description="For file operations developers usually use fs, for sqlite operations better-sqlite3 was used and for image operations I used sharp."
      />
      <FAQSection
        title="Some remarks regarding the Python implementation"
        description="For SQLite operations the Django ORM was used since this is the standard way to interact with SQL when choosing Python with Django. As for image processing, Pillow was chosen."
      />
    </div>
  </div>
</template>
