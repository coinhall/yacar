export const getRootDir = (): string => {
  const rootDir = process.env.ROOT_DIR;
  if (!rootDir) {
    console.error("Missing ROOT_DIR env variable");
    process.exit(1);
  }
  return rootDir;
};
