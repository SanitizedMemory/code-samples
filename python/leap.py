def foo(s, p):
    """
    Given a string s and a pattern p, return the indices of all occurrences of p in s.
    """
    return [idx for idx, _ in enumerate(s) if all(p.count(c) == s[idx: idx + len(p)].count(c) for c in set(p))]

