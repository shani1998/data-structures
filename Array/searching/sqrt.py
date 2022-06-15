# Returns floor of square root of x
def floorSqrt(x):
    # Base cases
    if x == 0 or x == 1:
        return x

        # Staring from 1, try all numbers until
    # i*i is greater than or equal to x.
    i = 1
    result = 1
    while result <= x:
        i += 1
        result = i * i

    return i - 1


def floorSqrt1(x):
    # Base cases
    if x == 0 or x == 1:
        return x
    # Do Binary Search for floor(sqrt(x))
    start = 1
    end = x
    ans = -1
    while start <= end:
        mid = (start + end) // 2
        # If x is a perfect square
        if mid * mid == x:
            return mid
        # answer when mid*mid is smaller
        # than x, and move closer to sqrt(x)
        if mid * mid < x:
            start = mid + 1
            ans = mid

        else:

            # If mid*mid is greater than x
            end = mid - 1

    return ans