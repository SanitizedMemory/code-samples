def is_anagram(one: str, two: str) -> bool:
    """ 
    This function determines if two strings are anagrams of each other.
    This is done by resolving 2 boolean expressions and combining their results.

    For the first boolean expression:
    we create a dictionary with a <char, bool> key/value pair. Each key is a unique
    character in the first string, and the value is `True` if that character
    appears in the second string the same number of times it appears in the first.
    If it does not appear the same amount of times, it is `False`. This is checked
    by subtracting the number of times it appears in each string and checking
    if it equals to zero. We then take this dictionary, and return True
    if all the values are True, and False if otherwise.

    For the second boolean expression:
    we check that both strings are of the same length. If there are any
    characters in the second word that do not appear in the first,
    it is possible for the first boolean expression to still return True.
    However, if this is the case, the second string must be longer than
    the first string; as it contains all the characters of the first string
    with the same number of occurences, plus some additional characters
    not seen in the first string.

    If both boolean expressions are True, then the strings are anagrams.
    Otherwise, they are not. The second boolean expression gets run first,
    but due to the nature of this logic it gets written last.

    Args:
        one (str): The first string in the comparison.
        two (str): The second string in the comparison.

    Returns:
        bool: True if the strings are anagrams, False otherwise.
    """
    return all({c: one.count(c) - two.count(c) == 0 for c in set(one)}.values()) if len(one) == len(two) else False

