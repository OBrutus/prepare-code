import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/**
 * Author  :   AV 
 * At      :   @Home 
 * Date    :   21.5.
 * ver.    :   0.0 
 * link    :   
 * file    :   cf/
 */

// public class I_AM_AV{
class I_AM_AV {
    static private I_AM_AV_Template sc = new I_AM_AV_Template();
    static private StringBuilder sb = new StringBuilder();

    public static int code(int[] a, int n) {
        return 0;
    }

    public static void main(String[] args) {
        int TEST_CASE = 0;
        int TOTAL_TEST_CASE = 1;

        TOTAL_TEST_CASE = sc.nextInt();
        
        while(TEST_CASE++ < TOTAL_TEST_CASE) {
            // int n=sc.nextInt();
            // String s = sc.next();
            // int[] a=sc.input(n);
            // var res = code(a, n);
            // System.out.println("Case #"+TEST_CASE+": "+res);
            // sb.append(res + "\n");
        }
        System.out.print(sb);
    }
}

class I_AM_AV_Template {
    static private BufferedReader br;
    static private StringTokenizer st;
    static public final int MOD = 1_000_000_007; // 10^9+7

    I_AM_AV_Template() {
        br = new BufferedReader(new InputStreamReader(System.in));
    }

    I_AM_AV_Template(String input) {
        try {
            // System.setIn(new FileInputStream(input));
            // br = new BufferedReader(new InputStreamReader(System.in));
            br = new BufferedReader(new InputStreamReader(new FileInputStream(input)));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    I_AM_AV_Template(String input, String output) {
        try {
            System.setIn(new FileInputStream(input));
            System.setOut(new PrintStream(output));
            br = new BufferedReader(new InputStreamReader(System.in));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * Use StringBuilder for output
     */
    String next() {
        while (st == null || !st.hasMoreElements()) {
            try {
                st = new StringTokenizer(br.readLine());
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        return st.nextToken();
    }

    int nextInt() {
        return Integer.parseInt(next());
    }

    long nextLong() {
        return Long.parseLong(next());
    }

    double nextDouble() {
        return Double.parseDouble(next());
    }

    String nextLine() {
        String str = "";
        try {
            str = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return str;
    }

    public int[] input(int n) {
        int[] a = new int[n];
        for (int i = 0; i < n; i++)
            a[i] = this.nextInt();
        return a;
    }

    public int[] input(int[] a, int n) {
        for (int i = 0; i < n; i++)
            a[i] = this.nextInt();
        return a;
    }

    public int[][] input(int r, int c){
        int[][] a = new int[r][c];
        for(int[] x : a) {
            input(x, c);
        }
        return a;
    }

    public void displayMatrix(int[][] a) {
        for (int[] x : a)
            System.out.println(Arrays.toString(x));
    }

    public int gcd(int a, int b) {
        if(b==0) return a;
        return gcd(b, a%b);
    }

    public boolean isPrime(int number) {
        if (number <= 2)
            return number == 2;
        else
            return (number % 2) != 0 && IntStream.rangeClosed(3, (int) Math.sqrt(number)).filter(n -> n % 2 != 0)
                    .noneMatch(n -> (number % n == 0));
    }
}
