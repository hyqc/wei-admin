/**
 * @see https://umijs.org/docs/max/access#access
 * */
export default function access(initialState: { currentUser?: API.CurrentUser } | undefined) {
  const { currentUser } = initialState ?? {};
  console.log(currentUser)
  return {
    // canAdmin: currentUser && currentUser.access === 'admin',
    allow: true,
    forbidden: false,
  };
}
