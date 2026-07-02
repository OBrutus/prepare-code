import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/**
 * Author       :   %s
 * Date         :   %s
 * ver.         :   %s
 * link         :   %s
 * dir          :   %s
 */

// public class OBrutusSolution {
class OBrutusSolution {
    private static final OBrutusSolutionTemplate sc = new OBrutusSolutionTemplate();
    private static final StringBuilder sb = new StringBuilder();

    public static int code(int[] a, int n) {
        return 0;
    }

    public static void main(String[] args) {
        // 64MB stack so deep DFS/recursion won't StackOverflow on big inputs.
        new Thread(null, OBrutusSolution::run, "main", 1 << 26).start();
    }

    private static void run() {
        // final int TOTAL_TEST_CASES = 1;
        final int TOTAL_TEST_CASES = sc.nextInt();

        for (int testCase = 1; testCase <= TOTAL_TEST_CASES; testCase++) {
            // Taking input
            // int n = sc.nextInt();
            // String s = sc.next();
            // int[] a = sc.input(n);

            // Compute
            var res = code(a, n);
            // System.out.println("DEBUG: Case #" + testCase + ": " + res);

            // Storing the result — chained append, NOT (res + "\n"): no per-line String garbage.
            sb.append(res).append('\n');
        }

        // Single bulk write of all output.
        try {
            System.out.write(sb.toString().getBytes());
            System.out.flush();
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }
}

class OBrutusSolutionTemplate {
    private final InputStream in;
    private final byte[] buf = new byte[1 << 16];
    private int ptr = 0, len = 0;
    public static final int MOD = 1_000_000_007; // 10^9+7

    OBrutusSolutionTemplate() {
        this.in = System.in;
    }

    OBrutusSolutionTemplate(String input) {
        InputStream s;
        try {
            s = new FileInputStream(input);
        } catch (IOException e) {
            e.printStackTrace();
            s = System.in;
        }
        this.in = s;
    }

    private int read() {
        if (ptr == len) {
            try {
                len = in.read(buf, 0, buf.length);
            } catch (IOException e) {
                throw new UncheckedIOException(e);
            }
            ptr = 0;
            if (len <= 0) return -1;
        }
        return buf[ptr++];
    }

    int nextInt() {
        int c = read();
        while (c <= ' ') {
            if (c == -1) throw new RuntimeException("EOF");
            c = read();
        }
        boolean neg = c == '-';
        if (neg) c = read();
        int x = 0;
        while (c >= '0' && c <= '9') {
            x = x * 10 + (c - '0');
            c = read();
        }
        return neg ? -x : x;
    }

    long nextLong() {
        int c = read();
        while (c <= ' ') {
            if (c == -1) throw new RuntimeException("EOF");
            c = read();
        }
        boolean neg = c == '-';
        if (neg) c = read();
        long x = 0;
        while (c >= '0' && c <= '9') {
            x = x * 10 + (c - '0');
            c = read();
        }
        return neg ? -x : x;
    }

    String next() {
        int c = read();
        while (c <= ' ') {
            if (c == -1) return null;
            c = read();
        }
        StringBuilder s = new StringBuilder();
        while (c > ' ') {
            s.append((char) c);
            c = read();
        }
        return s.toString();
    }

    double nextDouble() {
        return Double.parseDouble(next());
    }

    String nextLine() {
        StringBuilder s = new StringBuilder();
        int c = read();
        while (c != '\n' && c != -1) {
            if (c != '\r') s.append((char) c);
            c = read();
        }
        return (s.length() == 0 && c == -1) ? null : s.toString();
    }

    public int[] input(int n) {
        int[] a = new int[n];
        for (int i = 0; i < n; i++)
            a[i] = nextInt();
        return a;
    }

    public int[] input(int[] a, int n) {
        for (int i = 0; i < n; i++)
            a[i] = nextInt();
        return a;
    }

    public long[] inputLong(int n) {
        long[] a = new long[n];
        for (int i = 0; i < n; i++)
            a[i] = nextLong();
        return a;
    }

    public int[][] input(int r, int c) {
        int[][] a = new int[r][c];
        for (int[] x : a)
            input(x, c);
        return a;
    }

    public void displayMatrix(int[][] a) {
        for (int[] x : a)
            System.out.println(Arrays.toString(x));
    }

    public long gcd(long a, long b) {
        while (b != 0) {
            long t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    public boolean isPrime(int n) {
        if (n < 2) return false;
        if (n % 2 == 0) return n == 2;
        if (n % 3 == 0) return n == 3;
        for (int i = 5; (long) i * i <= n; i += 6)
            if (n % i == 0 || n % (i + 2) == 0) return false;
        return true;
    }
}
