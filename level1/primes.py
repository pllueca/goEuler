
def PrimeFactors(n):
    f = []
    i = 2
    while i*i <= n:
        if n%i :
             i += 1
        else:
            f.append(i)
            n //= i
    if n > i:
        f.append(n)
    return f
