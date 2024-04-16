module.exports = {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "type-enum": [
      2,
      "always",
      [
        "build",
        "chore",
        "ci",
        "docs",
        "feat",
        "fix",
        "perf",
        "refactor",
        "revert",
        "style",
        "test",
        "merge",
      ],
    ],
  },
  ignores: [
    function (commit) {
      return (
        commit.startsWith("Merge") ||
        commit.startsWith("Merged") ||
        commit.toLowerCase().startsWith("bump")
      );
    },
  ],
};
