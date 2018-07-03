export const formatTempVar = name =>
  `:${name.replace(/:/g, '').replace(/\s/g, '')}:`
