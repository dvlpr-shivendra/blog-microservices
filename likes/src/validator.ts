export function isValidPostId(postId: string): boolean {
  // Check if postId is not empty
  if (!postId || postId.trim() === "") {
    return false;
  }

  // If postIds are numeric, validate they only contain digits
  if (/^\d+$/.test(postId)) {
    return true;
  }

  return false;
}

export function isValidUserId(userId: string): boolean {
  // Check if userId is not empty
  if (!userId || userId.trim() === "") {
    return false;
  }

  // If userIds are numeric, validate they only contain digits
  if (/^\d+$/.test(userId)) {
    return true;
  }

  return false;
}
