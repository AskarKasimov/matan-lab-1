def sign(x):
    if x == 0:
        return 0
    return 1 if x > 0 else -1


def bisection_method(func, a, b, tol=0.0005):
    if func(a) == 0: return a, 0
    if func(b) == 0: return b, 0
    iters = 0
    while (b-a) > tol:
        iters += 1
        dx = (b-a)/2
        x_i = a+dx
        if sign(func(a)) != sign(func(x_i)): b = x_i
        else: a = x_i
    return x_i, tol, iters



def f(x):
    return pow(2, pow(x, 2))


res = bisection_method(f, 0, 1)
print(f"корень {res[0]} с точностью {res[1]}, кол-во итераций: {res[2]}")
