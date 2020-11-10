import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/**
 * Author :     AV
 * At   :       @Home
 * Date :       20.11.
 * ver. :       alpha-phase
 * link :       
 * file :       
 */

// public class I_AM_AV{
class I_AM_AV{
    static private I_AM_AV_Template sc=new I_AM_AV_Template();
    static private StringBuilder sb=new StringBuilder();
    static public HashMap<Integer, String> res = new HashMap<>();

    public static void main(String[] args) {
        final int TEST_CASES = sc.nextInt();
        // final int TEST_CASES = 1;
        for(int T=1; T<=TEST_CASES; T++){
            int n=sc.nextInt();
            int[] a=sc.input(n);
            // String s = sc.nextInt();
            new I_AM_AV_TestCase(T, a, n).start();
        }
        for(int i=1; i<=TEST_CASES; i++){
            System.out.printf("Case #%d%-3s", i, ":");
            System.out.println(res.get(i));
        }
    }

}

class I_AM_AV_TestCase extends Thread{
    int testCase;

    int[] a;
    int n;

    public I_AM_AV_TestCase(int testCase, int[] a, int n){
        this.testCase = testCase;
        this.a = a;
        this.n = n;
    }

    @Override
    public void run() {
        super.run();
        // Logic goes here !
        I_AM_AV.res.put(testCase, Arrays.stream(a).sum()+"");
    }

}

class I_AM_AV_Template{
    static private BufferedReader br;
    static private StringTokenizer st; 
    static public final int MOD=1000000007; //10^9+7

    
    I_AM_AV_Template(){
        br=new BufferedReader(new InputStreamReader(System.in));
    }  
    I_AM_AV_Template(String input){
        try {
            System.setIn(new FileInputStream(input));
            br=new BufferedReader(new InputStreamReader(System.in));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    I_AM_AV_Template(String input, String output){
        try {
            System.setIn(new FileInputStream(input));
            System.setOut(new PrintStream(output));
            br=new BufferedReader(new InputStreamReader(System.in));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    
    /**
     * Use StringBuilder for output
     */
    String next() { 
        while (st == null || !st.hasMoreElements()) { 
            try{ 
                st = new StringTokenizer(br.readLine()); 
            } 
            catch (IOException e) { 
                e.printStackTrace(); 
            } 
        } 
        return st.nextToken(); 
    } 
    int nextInt() { return Integer.parseInt(next()); } 
    long nextLong() { return Long.parseLong(next()); } 
    double nextDouble() {  return Double.parseDouble(next()); } 
    String nextLine() { 
        String str = ""; 
        try{ 
            str = br.readLine(); 
        } 
        catch (IOException e) { 
            e.printStackTrace(); 
        } 
        return str; 
    }
    public int[] input(int n){
        int[] a=new int[n];
        for(int i=0;i<n;i++)
            a[i]=this.nextInt();
        return a;
    }
    public int sum(int[] a, int start, int end){
        return Arrays.stream(a, start, end).sum();
    }
    public <E>void display(E[] a, int start, int end){
        for (;start<end; start++) {
            System.out.print(a[start]+" ");
        }
        System.out.println();
    }
    public void displayMatrix(int[][] a){
        for(int[] x: a)
            System.out.println(Arrays.toString(x));
    }

    public int max(int[] arr){ return Arrays.stream(arr).max().getAsInt(); }
    public int min(int[] arr){ return Arrays.stream(arr).min().getAsInt(); }

    public <E> E max(E[] a, int n){
        E max=a[0];
        return max;
    }
    public boolean isPrime(int number){
        if(number <= 2)
            return number == 2;
        else
            return  (number % 2) != 0
                    &&
                IntStream.rangeClosed(3, (int) Math.sqrt(number))
                    .filter(n -> n % 2 != 0)
                    .noneMatch(n -> (number % n == 0));
        }
}


